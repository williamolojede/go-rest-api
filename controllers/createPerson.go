package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = len(people) + 1
	fmt.Println(person, r.Body)
	people = append(people, person)
	json.NewEncoder(w).Encode(people)

}
