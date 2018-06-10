package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"encoding/json"
)

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, person := range people {
		// TODO:  handle error when isn't a number
		id, _ := strconv.Atoi(params["id"])

		if person.ID == id {
			json.NewEncoder(w).Encode(person)
			break
		}
	}
}
