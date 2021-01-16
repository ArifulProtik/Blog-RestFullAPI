package controller

import (
	"encoding/json"
	"forum-api/auth"
	"forum-api/db"
	"forum-api/utils"
	"io/ioutil"
	"log"
	"net/http"
)

// Signin Lets user sign in
func Signin(w http.ResponseWriter, r *http.Request) {
	type SigninModel struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6,max=20"`
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ErrorWriter(w, "Could not Process Any Reuest", http.StatusUnprocessableEntity)
		return
	}
	usr := SigninModel{}
	err = json.Unmarshal(body, &usr)
	if err != nil {
		utils.ErrorWriter(w, "Please Send Json", http.StatusUnprocessableEntity)
		return
	}
	if ok, errors := utils.Inputvalidate(usr); !ok {
		utils.JSONWriter(w, errors, http.StatusUnprocessableEntity)
		return
	}
	user, err := db.Dbprovider.FindUserByMail(usr.Email)
	if err != nil {
		utils.ErrorWriter(w, err.Error(), http.StatusNoContent)
		return
	}
	err = auth.VerifyPass(user.Password, usr.Password)
	if err != nil {
		utils.ErrorWriter(w, "Password Incorrect", 403)
		return
	}
	profile, err := db.Dbprovider.FindProfileByID(user.UID)
	if err != nil {
		utils.ErrorWriter(w, "Internal Error", http.StatusInternalServerError)
		return
	}
	token, err := auth.CreateToken(user.UID)
	if err != nil {
		utils.ErrorWriter(w, "Internal Error", http.StatusInternalServerError)
		return
	}
	tokenstring, err := db.Dbprovider.GetTokenByID(user.UID)
	log.Println(err)
	if err != nil {
		err = db.Dbprovider.SaveToken(user.UID, token)
		if err != nil {
			utils.ErrorWriter(w, err.Error(), http.StatusInternalServerError)
			return
		}
		utils.JSONWriter(w, data{
			"Token":   token,
			"Profile": profile,
		}, 200)
	} else {
		utils.JSONWriter(w, data{
			"Token":   tokenstring.Token,
			"Profile": profile,
		}, 200)
	}

}
