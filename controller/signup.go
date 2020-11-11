package controller

import (
	"encoding/json"
	"forum-api/auth"
	"forum-api/db"
	"forum-api/model"
	"forum-api/utils"
	"html"
	"io/ioutil"
	"net/http"
	"time"
)

// Signup let the user to create a Account.
func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	type recieverModel struct {
		FullName        string `json:"fullname" validate:"required"`
		Username        string `json:"username" validate:"required,min=6"`
		Email           string `json:"email" validate:"required,email"`
		Password        string `json:"password" validate:"required,min=6,max=20"`
		ConfirmPassword string `json:"confirmpassword" validate:"required"`
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ErrorWriter(w, "Could not Process Any Reuest", http.StatusUnprocessableEntity)
		return
	}
	user := recieverModel{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		utils.ErrorWriter(w, "Please Send Json", http.StatusUnprocessableEntity)
		return
	}
	if ok, errors := utils.Inputvalidate(user); !ok {
		utils.JSONWriter(w, errors, http.StatusUnprocessableEntity)
		return
	}
	if user.Password != user.ConfirmPassword {
		utils.ErrorWriter(w, "Password Does not match", http.StatusUnprocessableEntity)
		return
	}
	id := utils.UIDGen()
	Hashedpass, err := auth.HashBeforeSave(user.Password)
	if err != nil {
		utils.ErrorWriter(w, "Internal Error", http.StatusInternalServerError)
		return
	}
	usr := model.User{
		Email:     user.Email,
		Password:  string(Hashedpass),
		UID:       id,
		CreatedAt: time.Now(),
	}
	profile := model.Profile{
		Fullname:  html.EscapeString(user.FullName),
		Username:  html.EscapeString(user.Username),
		Email:     user.Email,
		UUID:      id,
		CreatedAt: time.Now(),
	}
	err = db.Dbprovider.AddUser(&usr)
	if err != nil {
		utils.ErrorWriter(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = db.Dbprovider.AddProfile(&profile)
	if err != nil {
		utils.ErrorWriter(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.SuccessWriter(w, "User Created Successfully", 200)
}
