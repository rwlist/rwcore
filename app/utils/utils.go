package utils

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error interface{}
}

type StdError struct {
	Code    int
	Message string
}

func WriteError(w http.ResponseWriter, code int, message string) {
	WriteJSON(
		w,
		code,
		Error{StdError{code, message}},
	)
}

func WriteJSON(w http.ResponseWriter, code int, body interface{}) {
	b, err := json.Marshal(body)
	if err != nil {
		// TODO: log error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(b)
}
