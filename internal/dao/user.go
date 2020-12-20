package dao

import (
	"gorm.io/gorm"

	"github.com/daheige/goapp/config"
	"github.com/daheige/goapp/internal/model"
)

// User user interface.
type User interface {
	GetUser(id int64) (*model.User, error)
}

type user struct {
	db *gorm.DB
}

// NewUserDao create an user dao.
func NewUserDao() User {
	return &user{
		db: config.GetDB("default"),
	}
}

// GetUser 根据id获取用户
func (u *user) GetUser(id int64) (*model.User, error) {
	var userInfo = &model.User{}
	err := u.db.Where("id = ?", id).First(userInfo).Error
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}
