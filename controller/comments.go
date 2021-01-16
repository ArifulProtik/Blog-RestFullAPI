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

	"github.com/gorilla/mux"
)

// GetComments gets All Comment Regrading Post
func GetComments(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["slug"]
	if params == "" {
		utils.ErrorWriter(w, "Missing Paramters", http.StatusUnprocessableEntity)
		return
	}
	cmts, err := db.Dbprovider.GetComments(params)
	if err != nil {
		utils.ErrorWriter(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	utils.JSONWriter(w, cmts, 200)
}

// SaveComment creates a comment
func SaveComment(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ErrorWriter(w, "Could not Process Any Reuest", http.StatusUnprocessableEntity)
		return
	}
	var comment model.Comment
	err = json.Unmarshal(body, &comment)
	if err != nil {
		utils.ErrorWriter(w, "Json Not Found", http.StatusUnprocessableEntity)
		return
	}
	if comment.Text == "" || comment.Postslug == "" {
		utils.ErrorWriter(w, "Empty Comment!!!!", http.StatusUnprocessableEntity)
		return
	}
	profile, err := auth.GetProfile(r)
	if err != nil {
		utils.ErrorWriter(w, "UnAuthorized", 403)
		return
	}
	newcomment := model.Comment{
		ID:         utils.UIDGen(),
		Authorname: profile.Fullname,
		Postslug:   comment.Postslug,
		Author:     profile,
		AuthorID:   profile.UUID,
		Text:       html.EscapeString(comment.Text),
	}
	added, err := db.Dbprovider.SaveComment(&newcomment)
	if err != nil {
		utils.ErrorWriter(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	utils.JSONWriter(w, added, 201)
}

// DeleteComment Deletes a Comment
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["uid"]
	if params == "" {
		utils.ErrorWriter(w, "No ID Provided", 404)
		return
	}
	cmt, err := db.Dbprovider.GetComment(params)
	if err != nil {
		utils.ErrorWriter(w, "No comment Found", 404)
		return
	}
	post, err := db.Dbprovider.Singlepost(cmt.Postslug)
	if err != nil {
		utils.ErrorWriter(w, "Unknown Post to Comment", 404)
		return
	}
	profile, err := auth.GetProfile(r)
	if err != nil {
		utils.ErrorWriter(w, "UnAuthorized", 403)
		return
	}
	if profile.UUID == cmt.Author.UUID || cmt.Author.UUID == post.AuthorID {
		err = db.Dbprovider.DeleteComment(cmt.ID)
		if err != nil {
			utils.ErrorWriter(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		utils.SuccessWriter(w, "Deleted!!!!!", 201)
		return
	}
	utils.ErrorWriter(w, "Unauthorized", 403)
}
