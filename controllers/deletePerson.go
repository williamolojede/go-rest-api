package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"encoding/json"
	. "github.com/williamolojede/rest-api/models"
)

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, person := range People {
		id, _ := strconv.Atoi(params["id"])
		if person.ID == id {
			People = append(People[:i], People[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(People)
}