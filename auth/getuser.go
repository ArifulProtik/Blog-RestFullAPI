package auth

import (
	"forum-api/db"
	"forum-api/model"
	"net/http"
)

// GetProfile get the User Profile
func GetProfile(r *http.Request) (model.Profile, error) {
	id, err := ExtractTokenID(r)
	if err != nil {
		return model.Profile{}, err
	}
	profile, err := db.Dbprovider.FindProfileByID(id)
	if err != nil {
		return model.Profile{}, err
	}
	return profile, nil

}
