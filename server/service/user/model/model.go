package model

type User struct {
	ID       int64  `gorm:"primaryKey"`
	Username string `gorm:"idx_username,unique"`
	Password string
}
