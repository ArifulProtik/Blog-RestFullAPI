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

	"github.com/gorilla/mux"
)

// GetAllPost Gets all Post
func GetAllPost(w http.ResponseWriter, r *http.Request) {
	posts, err := db.Dbprovider.Allpost()
	if err != nil {
		utils.ErrorWriter(w, "Internal Error", http.StatusInternalServerError)
		return
	}
	marshalled, err := utils.Postmarshal([]string{"list"}, posts)
	if err != nil {
		utils.ErrorWriter(w, "Internal Error", http.StatusInternalServerError)
		return
	}
	var marshalledposts []model.Post
	err = json.Unmarshal(marshalled, &marshalledposts)
	utils.JSONWriter(w, marshalledposts, 200)
}

// FeaturedPosts get all featured post
func FeaturedPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := db.Dbprovider.Fposts()
	if err != nil {
		utils.ErrorWriter(w, "Internal Error", http.StatusInternalServerError)
		return
	}
	marshalled, err := utils.Postmarshal([]string{"list"}, posts)
	if err != nil {
		utils.ErrorWriter(w, "Internal Error", http.StatusInternalServerError)
		return
	}
	var marshalledposts []model.Post
	err = json.Unmarshal(marshalled, &marshalledposts)
	utils.JSONWriter(w, marshalledposts, 200)
}

// CreatePost creates a post for Authenticated User
func CreatePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ErrorWriter(w, "Could not Process Any Reuest", http.StatusUnprocessableEntity)
		return
	}
	var post model.Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		utils.ErrorWriter(w, "Please Send Json", http.StatusUnprocessableEntity)
		return
	}
	if ok, errors := utils.Inputvalidate(post); !ok {
		utils.JSONWriter(w, errors, http.StatusUnprocessableEntity)
		return
	}
	userprofile, err := auth.GetProfile(r)
	if err != nil {
		utils.ErrorWriter(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// TODO: HTML Escape and GenSlug
	finalpost := model.Post{
		ID:           utils.UIDGen(),
		Title:        html.EscapeString(post.Title),
		Content:      html.EscapeString(post.Content),
		Author:       userprofile,
		Slug:         utils.GenSlug(post.Title),
		Featureimage: post.Featureimage,
		AuthorID:     userprofile.UUID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	newpost, err := db.Dbprovider.SavePost(&finalpost)
	if err != nil {
		utils.ErrorWriter(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.JSONWriter(w, newpost, 201)

}

// UpdatePost Updates a Post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ErrorWriter(w, "Could not Process Any Reuest", http.StatusUnprocessableEntity)
		return
	}
	var post model.Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		utils.ErrorWriter(w, "Please Send Json", http.StatusUnprocessableEntity)
		return
	}
	if ok, errors := utils.Inputvalidate(post); !ok {
		utils.JSONWriter(w, errors, http.StatusUnprocessableEntity)
		return
	}
	singlepost, err := db.Dbprovider.Singlepost(post.Slug)
	if err != nil {
		utils.ErrorWriter(w, "Internal Error", http.StatusInternalServerError)
		return
	}
	if singlepost.AuthorID != post.AuthorID {
		utils.ErrorWriter(w, "You have no Permission to edit this post", 403)
		return
	}
	updatedpost := model.Post{
		ID:           post.ID,
		Title:        html.EscapeString(post.Title),
		Content:      html.EscapeString(post.Content),
		Featureimage: post.Featureimage,
		UpdatedAt:    time.Now(),
	}
	newpost, err := db.Dbprovider.UpdatePost(&updatedpost)
	if err != nil {
		utils.ErrorWriter(w, "Internal Error", http.StatusInternalServerError)
		return
	}
	utils.JSONWriter(w, newpost, 201)

}

// Singlepost get a single post by slug
func Singlepost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["slug"]
	post, err := db.Dbprovider.Singlepost(params)
	if err != nil {
		utils.ErrorWriter(w, "Post Not Found", 404)
		return
	}
	utils.JSONWriter(w, post, 200)

}

// DeletePost Deletes a post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["slug"]
	post, err := db.Dbprovider.Singlepost(params)
	if err != nil {
		utils.ErrorWriter(w, "Post Not Found", 404)
		return
	}
	profile, err := auth.GetProfile(r)
	if err != nil {
		utils.ErrorWriter(w, "internal Error", http.StatusInternalServerError)
		return
	}
	if post.AuthorID != profile.UUID {
		utils.ErrorWriter(w, "You are unauthorized to Delete This post", 403)
		return
	}
	err = db.Dbprovider.DeletePost(post.Slug)
	if err != nil {
		utils.ErrorWriter(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	utils.SuccessWriter(w, "Post Deleted!!", 201)

}
