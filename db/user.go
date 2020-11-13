package db

import (
	"errors"
	"forum-api/model"

	"gorm.io/gorm"
)

func (c *manager) AddUser(usr *model.User) error {
	var err error
	err = c.db.Debug().Create(&usr).Error
	if err != nil {
		return err
	}
	return nil
}
func (c *manager) AddProfile(prfl *model.Profile) error {
	var err error
	err = c.db.Debug().Create(&prfl).Error
	if err != nil {
		return err
	}
	return nil
}
func (c *manager) FindUserByMail(mail string) (model.User, error) {
	var err error
	usr := model.User{}
	err = c.db.Debug().Model(model.User{}).Where("email = ?", mail).Take(&usr).Error
	if err != nil {
		return model.User{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.User{}, errors.New("No user found by the mail")
	}
	return usr, nil
}
func (c *manager) FindProfileByID(id string) (model.Profile, error) {
	var err error
	pfl := model.Profile{}
	err = c.db.Debug().Model(model.Profile{}).Where("uuid = ?", id).Take(&pfl).Error
	if err != nil {
		return model.Profile{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Profile{}, errors.New("No Profile Found By The ID")
	}
	return pfl, nil
}
