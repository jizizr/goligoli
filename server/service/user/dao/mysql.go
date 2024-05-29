package dao

import (
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/goligoli/server/common/consts"
	"github.com/jizizr/goligoli/server/service/user/model"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func (u *User) GetUserByName(username string) (*model.User, error) {
	var user model.User
	err := u.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) CreateUser(user *model.User) error {
	var temp model.User
	err := u.db.Where("username = ?", user.Username).First(&temp).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		klog.Errorf("MySQL error in GetUserByName: %v", err)
	}
	if temp.Username != "" {
		return errors.New(consts.MySqlAlreadyExist)
	}
	if err := u.db.Create(user).Error; err != nil {
		klog.Errorf("MySQL error in CreateUser: %v", err)
	}
	return nil
}

func NewUser(db *gorm.DB) *User {
	m := db.Migrator()
	if !m.HasTable(&model.User{}) {
		err := m.CreateTable(&model.User{})
		if err != nil {
			panic(err)
		}
	}
	return &User{db: db}
}
