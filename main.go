package main

import (
	"log"
	"net/http"
	"github.com/williamolojede/rest-api/routes"
	"github.com/williamolojede/rest-api/datastore"
)

var dao datastore.PeopleDAO

func init() {
	dao.DBURL = "localhost"
	dao.DBNAME = "people_db"
	dao.Connect()
}
func main() {

	log.Fatal(http.ListenAndServe(":8080", routes.RegisterRoutes()))
}
