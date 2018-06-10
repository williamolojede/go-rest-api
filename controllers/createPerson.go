package controllers

import (
	"net/http"
	"encoding/json"
	. "github.com/williamolojede/rest-api/models"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = len(People) + 1
	People = append(People, person)
	json.NewEncoder(w).Encode(People)

}
