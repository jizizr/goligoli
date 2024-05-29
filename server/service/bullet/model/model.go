package model

import "time"

type Bullet struct {
	ID       int64 `gorm:"primaryKey"`
	UserID   int64
	LiveID   int64
	LiveTime int64     `gorm:"idx_live_time"`
	SendTime time.Time `gorm:"autoCreateTime"`
	Content  string
}
