package db

import (
	"errors"
	"forum-api/model"

	"gorm.io/gorm"
)

func (c *manager) SavePost(p *model.Post) (*model.Post, error) {
	newpost := c.db.Debug().Create(&p)
	if newpost.Error != nil {
		return p, newpost.Error
	}
	return p, nil

}
func (c *manager) UpdatePost(p *model.Post) (*model.Post, error) {
	updatedpost := c.db.Save(&p)
	if updatedpost.Error != nil {
		return p, updatedpost.Error
	}
	return p, nil
}
func (c *manager) Allpost() ([]model.Post, error) {
	var posts []model.Post
	getposts := c.db.Limit(20).Find(&posts)
	if getposts.Error != nil {
		return []model.Post{}, getposts.Error
	}
	return posts, nil
}
func (c *manager) GetComments(slug string) ([]model.Comment, error) {
	var comments []model.Comment
	getcomments := c.db.Limit(20).Find(&comments)
	if getcomments.Error != nil {
		return []model.Comment{}, getcomments.Error
	}
	return comments, nil
}

func (c *manager) Singlepost(postslug string) (model.Post, error) {
	var post model.Post
	getpost := c.db.Where("slug=?", postslug).Find(&post)
	if getpost.Error != nil {
		return model.Post{}, getpost.Error
	}
	return post, nil

}
func (c *manager) GetComment(UUID string) (model.Comment, error) {
	var cmt model.Comment
	getcmt := c.db.Where("id=?", UUID).Find(&cmt)
	if getcmt.Error != nil {
		return model.Comment{}, getcmt.Error
	}
	return cmt, nil
}
func (c *manager) Postbyuser(userid string) ([]model.Post, error) {
	var posts []model.Post
	getposts := c.db.Where("authorid=?", userid).Find(&posts)
	if getposts.Error != nil {
		return []model.Post{}, getposts.Error
	}
	return posts, nil

}
func (c *manager) DeletePost(postid string) error {
	var post model.Post
	err := c.db.Where("id=?", postid).Find(&post)
	if err.Error != nil {
		return err.Error
	}
	c.db.Delete(&post)
	return nil
}
func (c *manager) SaveComment(comment *model.Comment) (*model.Comment, error) {
	newcomment := c.db.Create(&comment)
	if newcomment.Error != nil {
		return &model.Comment{}, newcomment.Error
	}
	return comment, nil
}
func (c *manager) DeleteComment(cid string) error {
	var comment model.Comment
	err := c.db.Where("id=?", comment).Find(&comment)
	if err.Error != nil {
		return err.Error
	}
	c.db.Delete(&comment)
	return nil
}
func (c *manager) DoLike(like *model.Like, pid string) error {
	err := c.db.Where("userid=?", pid).Find(&like).Error
	rec := errors.Is(err, gorm.ErrRecordNotFound)
	if !rec {
		c.db.Delete(&like)
		return nil
	}
	c.db.Create(&like)
	return nil

}
