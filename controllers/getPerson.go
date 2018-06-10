package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/williamolojede/rest-api/helpers"
)

func GetPerson(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	person, err := dao.FindById(params["id"])
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid person ID")
		return
	}
	helpers.RespondWithJson(w, http.StatusOK, person)
}
