package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/williamolojede/rest-api/helpers"
)

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := dao.Delete(params["id"]); err != nil {
		helpers.RespondWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	helpers.RespondWithJson(w, http.StatusOK, map[string]string{"result": "person deleted"})

}