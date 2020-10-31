package db

import "forum-api/model"

func (c *manager) AddUser(usr *model.User) error {
	var err error
	err := c.db.Debug().Create(&usr).Error
	if err != nil {
		return err
	}
	return nil
}
func (c *manager) AddProfile(prfl *model.Profile) error {
	var err error
	err := c.db.Debug().Create(&prfl).Error
	if err != nil {
		return err
	}
	return nil
}
