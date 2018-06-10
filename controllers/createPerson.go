package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/williamolojede/rest-api/datastore"
	"github.com/williamolojede/rest-api/helpers"
	"github.com/globalsign/mgo/bson"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var person datastore.Person

	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	person.ID = bson.NewObjectId()


	if err := dao.Insert(person); err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.RespondWithJson(w, http.StatusCreated, person)
}
