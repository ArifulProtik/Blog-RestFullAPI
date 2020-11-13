package db

import (
	"errors"
	"forum-api/model"
	"time"

	"gorm.io/gorm"
)

func (c *manager) SaveToken(id string, token string) error {
	var err error
	var acess model.TokenString
	acess = model.TokenString{
		UUID:   id,
		Token:  token,
		Expire: time.Now().Add(280 * time.Hour),
	}
	err = c.db.Debug().Create(&acess).Error
	if err != nil {
		return err
	}
	return nil
}
func (c *manager) GetTokenByID(id string) (model.TokenString, error) {
	var err error
	acess := model.TokenString{}
	err = c.db.Debug().Model(model.TokenString{}).Where("uuid = ?", id).Take(&acess).Error
	if err != nil {
		return model.TokenString{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.TokenString{}, errors.New("No token Found")
	}
	return acess, nil
}
func (c *manager) DeleteToken(token string) (int64, error) {
	err := c.db.Debug().Model(&model.TokenString{}).Where("token = ?", token).Take(&model.TokenString{}).Delete(&model.TokenString{}).Error
	if err != nil {
		return 0, err
	}
	return c.db.RowsAffected, nil
}
