package dao

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/goligoli/server/kitex_gen/base"
	"gorm.io/gorm"
)

type Live struct {
	db *gorm.DB
}

func (l *Live) AddLiveRoom(room *base.Room) (int64, error) {
	var temp base.Room
	err := l.db.Where("owner = ?", room.Owner).Find(&temp).Error

	if err != gorm.ErrRecordNotFound && err != nil {
		klog.Errorf("MySql Error in AddLiveRoom: %v", err)
		return temp.LiveId, err
	}

	if temp.LiveId != 0 {
		return temp.LiveId, nil
	}

	if err := l.db.Create(room).Error; err != nil {
		klog.Errorf("CreateLiveRoom failed: %v", err)
	}
	return temp.LiveId, nil
}

func (l *Live) DeleteLiveRoom(id int64) error {
	err := l.db.Where("live_id = ?", id).Delete(&base.Room{}).Error
	if err != nil {
		klog.Errorf("DeleteLiveRoom failed: %v", err)
		return err
	}
	return nil
}

func (l *Live) GetLiveRoomByID(id int64) (*base.Room, error) {
	var room base.Room
	err := l.db.Where("live_id = ?", id).Find(&room).Error
	if err == gorm.ErrRecordNotFound {
		return &room, nil
	} else if err != nil {
		klog.Errorf("query liveroom failed: %v", err)
		return nil, err
	}
	return &room, err
}

func (l *Live) GetLiveRoomOwnerByID(id int64) (int64, error) {
	var owner int64
	err := l.db.Model(&base.Room{}).Where("live_id = ?", id).Select("owner").First(&owner).Error
	if err == gorm.ErrRecordNotFound {
		return 0, nil
	} else if err != nil {
		klog.Errorf("query liveroom owner failed: %v", err)
		return 0, err
	}
	return owner, err
}

func NewLive(db *gorm.DB) *Live {
	m := db.Migrator()
	if !m.HasTable(&base.Room{}) {
		err := m.CreateTable(&base.Room{})
		if err != nil {
			panic(err)
		}
	}
	return &Live{db}
}
