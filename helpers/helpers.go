package helpers

import (
	"net/http"
	"encoding/json"
)

//func RespondWithError(w http.ResponseWriter, code int, msg string) {
//	response, _ := json.Marshal(map[string]string{"error": msg})
//	w.WriteHeader(code)
//	w.Write(response)
//}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJson(w, code, map[string]string{"error": msg})
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	// data, err := json.MarshalIndent(v, "", "  ")
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(response)
}