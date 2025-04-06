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

func writeResponseError(w http.ResponseWriter, message string, code int) {
	resp := Response{
		Message: message,
		Code:    code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

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
	StatusOkHandler = func(w http.ResponseWriter, message string) {
		writeResponse(w, message, http.StatusOK)
	}
	RequestErrorHandler = func(w http.ResponseWriter, message string, code int) {
		writeResponseError(w, message, code)
	}
	RequestSuccessHandler = func(w http.ResponseWriter, message string, code int) {
		writeResponse(w, message, code)
	}
)
