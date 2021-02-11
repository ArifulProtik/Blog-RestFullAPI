package model

import "time"

// Post For blog post
type Post struct {
	ID           string    `gorm:"primary_key;" json:"id" groups:"list,single"`
	Title        string    `gorm:"size:255;not null;unique" json:"title" validate:"required,min=6" groups:"list,single"`
	Content      string    `sql:"type:text"  gorm:"not null;" json:"content" validate:"required" groups:"single"`
	Author       Profile   `gorm:"foreignKey:AuthorID" json:"author" groups:"list,single"`
	Slug         string    `gorm:"size:255, not null" json:"slug" groups:"list,single"`
	Featureimage string    `gorm:"not null" json:"f_image" groups:"list,single"`
	AuthorID     string    `gorm:"not null" json:"author_id" groups:"list,single"`
	Featured     bool      `gorm:"default:false" json:"featured"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at" groups:"list,single"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at" groups:"list,single"`
}

// Comment for comment on a Single post
type Comment struct {
	ID         string  `gorm:"primary_key;" json:"Id"`
	Authorname string  `gorm:"size:255" json:"Name"`
	Postslug   string  `json:"Pslug"`
	Author     Profile `gorm:"foreignKey:AuthorID" json:"author"`
	AuthorID   string  `json:"authorid"`
	Text       string  `json:"Text" validate:"required,min=6"`
}

// Like Hold The post like Info
type Like struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserID    string    `gorm:"not null" json:"user_id"`
	PostID    string    `gorm:"not null" json:"post_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
