package model

import "time"

// TokenString for validation and Expiracy of Aceesss Token
type TokenString struct {
	UUID   string `gorm:"primary_key;unique,not null"`
	Token  string `gorm:"unique"`
	Expire time.Time
}
