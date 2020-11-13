package controller

import "net/http"

//TODO CORS middleware

// JSONandCORS a Middleware to Setup CORS and JSON
func JSONandCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json; charset=UTF8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r)
	}
}

//TODO Log Middleware
//TODO AUth Middleware
