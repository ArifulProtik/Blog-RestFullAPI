package model

import "time"

// User hold the internal UserInformation
type User struct {
	Email     string    `gorm:"size:255; unique" json:"email"`
	Password  string    `gorm:"size:255" json:"password"`
	UID       string    `gorm:"primary_key; unique" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	IsActive  bool      `gorm:"default:false" json:"isactive"`
	Role      string    `gorm:"default:false" json:"role"`
}

// Profile holds Public Userinformation
type Profile struct {
	Fullname     string    `gorm:"size:255" json:"fullname"`
	Email        string    `gorm:"size:255" json:"email"`
	Profileimage string    `gorm:"size:255" json:"profile_pic"`
	UUID         string    `gorm:"primary_key; unique" json:"id"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Username     string    `gorm:"size:255; unique" json:"username"`
	sociallink
}
type sociallink struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
	Github   string `json:"github"`
	Linkdin  string `json:"linkdin"`
}
