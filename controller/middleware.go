package controller

import (
	"forum-api/auth"
	"forum-api/db"
	"forum-api/utils"
	"net/http"
)

//TODO CORS middleware

// JSONandCORS a Middleware to Setup CORS and JSON
func JSONandCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json; charset=UTF8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r)
	}
}
 
// IsAuth Authentication Check Middleware.
func IsAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			utils.ErrorWriter(w, "You are Unauthorized!!!!!!!!", 403)
			return
		}
		id, err := auth.ExtractTokenID(r)
		if err != nil {
			utils.ErrorWriter(w, "Internal Error", http.StatusInternalServerError)
			return
		}
		_, err = db.Dbprovider.GetTokenByID(id)
		if err != nil {
			utils.ErrorWriter(w,"Unauthorized", 403)
			return
		}
		err = auth.TokenValid(r)
		if err != nil {
			utils.ErrorWriter(w, "Unauthoroized", 403)
			return
		}
		// TODO Check  Expiration of token For Further Validation 
		next(w,r)
	}
}
