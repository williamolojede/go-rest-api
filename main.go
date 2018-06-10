package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/williamolojede/rest-api/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
