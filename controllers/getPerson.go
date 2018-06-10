package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"encoding/json"
	. "github.com/williamolojede/rest-api/models"
)

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, person := range People {
		// TODO:  handle error when isn't a number
		id, _ := strconv.Atoi(params["id"])

		if person.ID == id {
			json.NewEncoder(w).Encode(person)
			break
		}
	}
}
