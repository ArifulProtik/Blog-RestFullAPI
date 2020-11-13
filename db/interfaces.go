package db

import "forum-api/model"

// Manager Defines all Database Operation.
type Manager interface {
	PingDB()

	AddUser(*model.User) error                     // Insert User Data on DB
	AddProfile(*model.Profile) error               // Insert profile on DB
	FindUserByMail(string) (model.User, error)     // Find User By Mail
	FindProfileByID(string) (model.Profile, error) // Find Profile BY ID

	SaveToken(string, string) error                 // Saves The Session to DB
	GetTokenByID(string) (model.TokenString, error) // gets the existing session
	DeleteToken(string) (int64, error)              // Destroy The session
}
