package controllers

import (
	"net/http"
	"github.com/williamolojede/rest-api/helpers"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	people, err := dao.FindAll()
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helpers.RespondWithJson(w, http.StatusOK, people)
}
