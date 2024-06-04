package dao

import (
	"github.com/jizizr/goligoli/server/kitex_gen/base"
	"gorm.io/gorm"
)

type Message struct {
	db *gorm.DB
}

func (b *Message) CreateMessage(Message *base.LiveMessage) error {
	err := b.db.Create(Message).Error
	if err != nil {
		return err
	}
	return nil
}

func (b *Message) GetMessageByID(id int64) (*base.LiveMessage, error) {
	var Message base.LiveMessage
	err := b.db.Where("id = ?", id).First(&Message).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &Message, nil
}

func (b *Message) GetHistoryMessagesByTime(liveID int64, startTime int64, offset int64) ([]*base.LiveMessage, error) {
	var Messages []*base.LiveMessage
	err := b.db.Where("live_id = ? AND live_time BETWEEN ? AND ?", liveID, startTime, startTime+offset).Order("live_time").Limit(100).Find(&Messages).Error
	if err != nil {
		return nil, err
	}
	return Messages, nil
}

func NewMessage(db *gorm.DB) *Message {
	m := db.Migrator()
	if !m.HasTable(&base.LiveMessage{}) {
		err := m.CreateTable(&base.LiveMessage{})
		if err != nil {
			panic(err)
		}
	}
	return &Message{db: db}
}
