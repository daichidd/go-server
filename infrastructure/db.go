package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

var DB *sql.DB

func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		return errors.Wrap(err, "dotenv.Load failed")
	}
	return nil
}

func init() {
	// load dotenv
	if err := loadEnv(); err != nil {
		log.Fatal(err)
	}

	driverName := os.Getenv("DRIVER_NAME")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_CONTAINER_NAME")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	protocol := os.Getenv("MYSQL_PROTOCOL")
	port := os.Getenv("MYSQL_PORT")

	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true", user, pass, protocol, host, port, dbName)

	var err error
	DB, err = sql.Open(driverName, dsn)
	if err != nil {
		log.Fatal(err, "sql.Open failed")
	}
}
