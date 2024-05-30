package model

type User struct {
	ID       int64  `gorm:"type:bigint;primaryKey;autoIncrement:false"`
	Username string `gorm:"type:varchar(32);uniqueIndex"`
	Password string `gorm:"type:varchar(32)"`
}
