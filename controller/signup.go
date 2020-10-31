package controller

import "net/http"

// Signup let the user to create a Account.
func Signup(w http.ResponseWriter, r *http.Request) {
	type recieverModel struct {
		FullName        string `json:"fullname"`
		Username        string `json:"username"`
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmpassword"`
	}
}
