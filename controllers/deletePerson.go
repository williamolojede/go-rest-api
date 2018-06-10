package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"encoding/json"
)

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, person := range people {
		id, _ := strconv.Atoi(params["id"])
		if person.ID == id {
			people = append(people[:i], people[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}