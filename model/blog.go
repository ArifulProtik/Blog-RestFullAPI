package model

import "time"

// Post For blog post
type Post struct {
	ID        string    `gorm:"primary_key;" json:"id"`
	Title     string    `gorm:"size:255;not null;unique" json:"title"`
	Content   string    `gorm:"not null;" json:"content"`
	Author    User      `json:"author"`
	Slug      string    `gorm:"size:255, not null" json:"slug"`
	AuthorID  string    `gorm:"not null" json:"author_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Comment for comment on a Single post
type Comment struct {
	ID         string `gorm:"primary_key;" json:"Id"`
	Authorname string `gorm:"size:255" json:"Name"`
	Postid     string `json:"Pid"`
	Author     User   `json:"author"`
	Text       string `json:"Text"`
}
// Like Hold The post like Info
type Like struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserID    string    `gorm:"not null" json:"user_id"`
	PostID    string    `gorm:"not null" json:"post_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}