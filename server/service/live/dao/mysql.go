package dao

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/goligoli/server/kitex_gen/base"
	"gorm.io/gorm"
)

type Live struct {
	db *gorm.DB
}

func (l *Live) GetOnlineLiveRooms() ([]int64, error) {
	var rooms []int64
	err := l.db.Model(&base.Room{}).Where("is_live = 1").Pluck("live_id", &rooms).Error
	if err != nil {
		klog.Errorf("GetOnlineLiveRooms failed: %v", err)
	}
	return rooms, err
}

func (l *Live) AddLiveRoom(room *base.Room) (int64, error) {
	var temp base.Room
	err := l.db.Where("owner = ? AND is_live = 1", room.Owner).Find(&temp).Error

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

func (l *Live) TagStopLiveRoom(id int64) error {
	err := l.db.Model(&base.Room{}).Where("live_id = ?", id).Update("is_live", false).Error
	if err != nil {
		klog.Errorf("TagStopLiveRoom failed: %v", err)
	}
	return err
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
