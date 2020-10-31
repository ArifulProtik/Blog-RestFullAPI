package controller

import (
	"forum-api/utils"
	"net/http"
)

type data map[string]interface{}

// Home Writes the Home Response
func Home(w http.ResponseWriter, r *http.Request) {
	utils.JSONWriter(w, data{
		"Messege": "Welcome To Forum",
	}, 200)
}
