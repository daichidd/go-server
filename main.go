package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DaichiEndo/default/handler"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Println("hello")
	})
	router.GET("/dbHealthCheck", handler.NewDB().DBHealthCheck)

	router.GET("/user", handler.NewUser().Lister)
	router.POST("/setUser", handler.NewUser().Setter)

	log.Fatal(http.ListenAndServe(":8080", router))
}
