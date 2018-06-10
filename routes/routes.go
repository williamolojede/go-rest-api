package routes

import (
	"github.com/gorilla/mux"
	"github.com/williamolojede/rest-api/controllers"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/people", controllers.GetPeople).Methods("GET")
	r.HandleFunc("/people/{id}", controllers.GetPerson).Methods("GET")
	r.HandleFunc("/people", controllers.CreatePerson).Methods("POST")
	r.HandleFunc("/people/{id}", controllers.DeletePerson).Methods("DELETE")
	return r
}


