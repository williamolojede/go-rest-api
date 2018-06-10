package routes

import (
	"github.com/gorilla/mux"
	"github.com/williamolojede/rest-api/controllers"
	"github.com/williamolojede/rest-api/middlewares"
)

func RegisterRoutes(r *mux.Router) {
	r.Use(middlewares.JsonResponse)
	r.HandleFunc("/people", controllers.GetPeople).Methods("GET")
	r.HandleFunc("/people/{id}", controllers.GetPerson).Methods("GET")
	r.HandleFunc("/people", controllers.CreatePerson).Methods("POST")
	r.HandleFunc("/people/{id}", controllers.DeletePerson).Methods("DELETE")
}

