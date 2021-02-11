package db

import "forum-api/model"

// Manager Defines all Database Operation.
type Manager interface {
	PingDB()

	AddUser(*model.User) error                     // Insert User Data on DB
	AddProfile(*model.Profile) error               // Insert profile on DB
	FindUserByMail(string) (model.User, error)     // Find User By Mail
	FindProfileByID(string) (model.Profile, error) // Find Profile BY ID
	// Stored Session

	SaveToken(string, string) error                 // Saves The Session to DB
	GetTokenByID(string) (model.TokenString, error) // gets the existing session
	DeleteToken(string) (int64, error)              // Destroy The session
	// Post
	SavePost(*model.Post) (*model.Post, error)   // Creates a post
	UpdatePost(*model.Post) (*model.Post, error) // Edits a post
	Allpost() ([]model.Post, error)              // Getall post
	Fposts() ([]model.Post, error)               // Featured Post only ^
	Singlepost(string) (model.Post, error)       // Gets a Single post
	Postbyuser(string) ([]model.Post, error)     // Get all post by single user
	DeletePost(string) error                     // Delete single post
	// Comment
	SaveComment(*model.Comment) (*model.Comment, error) // Creates Comment
	DeleteComment(string) error                         // Deletes a Comment
	GetComment(string) (model.Comment, error)           // Gets A Soingle Comment
	GetComments(string) ([]model.Comment, error)
	// like
	DoLike(*model.Like, string) error // Likes or Dislikes a Post
}
