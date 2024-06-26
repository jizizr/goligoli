package dao

import (
	"github.com/cloudwego/kitex/pkg/klog"
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

func (u *User) CreateUser(user *model.User) (string, error) {
	var temp model.User
	err := u.db.Where("username = ?", user.Username).First(&temp).Error
	// if err isn't nil and RecordNotFound return error
	if err != gorm.ErrRecordNotFound && err != nil {
		klog.Errorf("MySQL error in GetUserByName: %v", err)
		return temp.Username, err
	}
	// if Username isn't empty, return username to indicate that the user already exists
	if temp.Username != "" {
		return temp.Username, nil
	}

	// if Username is empty, it's a new user, create it
	if err := u.db.Create(user).Error; err != nil {
		klog.Errorf("MySQL error in CreateUser: %v", err)
	}
	return "", nil
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
