package helpers

import (
	"testing"
	"net/http"
	"encoding/json"
)

var actualPayload []byte
var actualStatusCode int

func TestRespondWithJson(t *testing.T) {
	var responseWriter mockResponseWriter
	responseWriter.header = make(map[string][]string)
	payload := map[string]string{"message": "person updated"}

	expectedContentType := "application/json; charset=utf-8"
	expectedStatusCode := http.StatusOK
	expectedPayload, _ := json.MarshalIndent(payload, "", "  ")

	RespondWithJson(responseWriter, expectedStatusCode, payload)

	actualContentType := responseWriter.Header().Get("Content-Type")

	if expectedContentType != actualContentType{
		t.Errorf("expected content-type to be '%v' but got '%v'", expectedContentType, actualContentType)
	}

	if expectedStatusCode != actualStatusCode {
		t.Errorf("expected statusCode to be '%v' but got '%v'", expectedStatusCode, actualStatusCode)
	}

	if string(expectedPayload) != string(actualPayload) {
		t.Error("Incorrect payload sent to client")
	}
}

func TestRespondWithError(t *testing.T) {
	var responseWriter mockResponseWriter
	responseWriter.header = make(map[string][]string)
	payload := map[string]string{"message": "user not found"}

	expectedStatusCode := http.StatusNotFound
	expectedPayload, _ := json.MarshalIndent(payload, "", "  ")

	RespondWithError(responseWriter, expectedStatusCode, "user not found")

	if expectedStatusCode != actualStatusCode {
		t.Errorf("expected statusCode to be '%v' but got '%v'", expectedStatusCode, actualStatusCode)
	}

	if string(expectedPayload) != string(actualPayload) {
		t.Error("Incorrect payload sent to client")
	}
}

type mockResponseWriter struct {
	header http.Header
}

func (mrw mockResponseWriter) Header() http.Header {
	return mrw.header
}
func (mrw mockResponseWriter) Write(data []byte) (int, error) {
	actualPayload = data
	return len(data), nil
}

func (mrw mockResponseWriter) WriteHeader(statusCode int) {
	actualStatusCode = statusCode
}