package write_response

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string
	Code    int
}

// writeResponseError is used to handle error responses in a standardized format.
func writeResponseError(w http.ResponseWriter, message string, code int) {
	resp := Response{
		Message: message,
		Code:    code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

// writeResponse is used to handle successful responses in a standardized format.
func writeResponse(w http.ResponseWriter, message string, code int) {
	resp := Response{
		Message: message,
		Code:    code,
	}

	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonResponse)
}

var (
	// StatusOkHandler is a reusable handler to send a success message with HTTP status 200 (OK).
	StatusOkHandler = func(w http.ResponseWriter, message string) {
		writeResponse(w, message, http.StatusOK)
	}
	// RequestErrorHandler is a reusable handler to send an error message with a given HTTP status code.
	RequestErrorHandler = func(w http.ResponseWriter, message string, code int) {
		writeResponseError(w, message, code)
	}
	// RequestSuccessHandler is a reusable handler to send a success message with a given HTTP status code.
	RequestSuccessHandler = func(w http.ResponseWriter, message string, code int) {
		writeResponse(w, message, code)
	}
)
