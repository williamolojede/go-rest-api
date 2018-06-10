package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/williamolojede/rest-api/datastore"
	"github.com/williamolojede/rest-api/helpers"
)

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var person datastore.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(person); err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helpers.RespondWithJson(w, http.StatusOK, map[string]string{"result": "person updated"})
}