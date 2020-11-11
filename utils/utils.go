package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type text map[string]interface{}

// JSONWriter Writes the Json to the Response
func JSONWriter(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}

// ErrorWriter writes error to the Response
func ErrorWriter(w http.ResponseWriter, msg string, statusCode int) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(text{
		"Error": msg,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}

// SuccessWriter Writes Success to the Response
func SuccessWriter(w http.ResponseWriter, msg string, statusCode int) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(text{
		"Success": msg,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}
