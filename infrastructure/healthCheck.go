package infrastructure

import (
	"database/sql"

	"github.com/pkg/errors"
)

type DBHealthCheckRepository struct {
	db *sql.DB
}

func NewDBHealthCheck(db *sql.DB) *DBHealthCheckRepository {
	return &DBHealthCheckRepository{
		db: db,
	}
}

func (r *DBHealthCheckRepository) DBHealthCheck() (string, error) {
	if err := r.db.Ping(); err != nil {
		return "", errors.Wrap(err, "failed db.Ping")
	}
	return "Pong", nil
}
