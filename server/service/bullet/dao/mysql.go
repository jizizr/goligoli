package dao

import (
	"github.com/jizizr/goligoli/server/kitex_gen/base"
	"gorm.io/gorm"
)

type Bullet struct {
	db *gorm.DB
}

func (b *Bullet) CreateBullet(bullet *base.Bullet) error {
	err := b.db.Create(bullet).Error
	if err != nil {
		return err
	}
	return nil
}

func (b *Bullet) GetBulletByID(id int64) (*base.Bullet, error) {
	var bullet base.Bullet
	err := b.db.Where("id = ?", id).First(&bullet).Error
	if err != nil {
		return nil, err
	}
	return &bullet, nil
}

func (b *Bullet) GetHistoryBulletsByTime(liveID int64, startTime int64, offset int64) ([]*base.Bullet, error) {
	var bullets []*base.Bullet
	err := b.db.Where("live_id = ? AND live_time BETWEEN ? AND ?", liveID, startTime, startTime+offset).Order("create_time").Limit(100).Find(&bullets).Error
	if err != nil {
		return nil, err
	}
	return bullets, nil
}

func NewBullet(db *gorm.DB) *Bullet {
	m := db.Migrator()
	if !m.HasTable(&base.Bullet{}) {
		err := m.CreateTable(&base.Bullet{})
		if err != nil {
			panic(err)
		}
	}
	return &Bullet{db: db}
}
