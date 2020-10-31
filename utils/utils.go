package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSONWriter Writes the Json to the Response
func JSONWriter(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}
