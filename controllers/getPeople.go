package controllers

import (
	"net/http"
	"encoding/json"
	. "github.com/williamolojede/rest-api/models"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(People)
}
