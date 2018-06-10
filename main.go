package main

import (
	"log"
	"net/http"
	"github.com/williamolojede/rest-api/routes"
)

func main() {

	log.Fatal(http.ListenAndServe(":8080", routes.RegisterRoutes()))
}
