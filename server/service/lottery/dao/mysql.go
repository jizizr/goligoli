package dao

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/goligoli/server/kitex_gen/base"
	"github.com/jizizr/goligoli/server/service/lottery/model"
	"gorm.io/gorm"
)

type Winner struct {
	db *gorm.DB
}

func (l *Winner) AddWinners(lotteryId int64, users []int64) error {
	err := l.db.Create(&model.Result{
		ID:      lotteryId,
		Winners: users,
	}).Error
	if err != nil {
		klog.Errorf("failed to add winners: %v", err)
		return err
	}
	return nil
}

func (l *Winner) GetWinners(lotteryId int64) ([]int64, error) {
	var result model.Result
	err := l.db.Where("id = ?", lotteryId).First(&result).Error
	if err != nil {
		klog.Errorf("failed to get winners: %v", err)
		return nil, err
	}
	return result.Winners, nil
}

func NewWinner(db *gorm.DB) *Winner {
	m := db.Migrator()
	if !m.HasTable(&model.Result{}) {
		err := m.CreateTable(&model.Result{})
		if err != nil {
			panic(err)
		}
	}
	return &Winner{db: db}
}

type Lottery struct {
	db *gorm.DB
}

func (l *Lottery) SetLotteryDB(info *base.Gift) error {
	err := l.db.Create(info).Error
	if err != nil {
		klog.Errorf("failed to create lottery: %v", err)
		return err
	}
	return nil
}

func (l *Lottery) GetLotteryDB(LiveID int64) ([]*base.Gift, error) {
	var records []*base.Gift
	err := l.db.Where("id = ?", LiveID).Find(&records).Error
	if err != nil {
		klog.Errorf("failed to get lottery: %v", err)
		return nil, err
	}
	return records, nil
}

func (l *Lottery) GetLotteryByID(lotteryID int64) (*base.Gift, error) {
	var record base.Gift
	err := l.db.Where("id = ?", lotteryID).First(&record).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		klog.Errorf("failed to get lottery: %v", err)
		return nil, err
	}
	return &record, nil

}

func NewLottery(db *gorm.DB) *Lottery {
	m := db.Migrator()
	if !m.HasTable(&base.Gift{}) {
		err := m.CreateTable(&base.Gift{})
		if err != nil {
			panic(err)
		}
	}
	return &Lottery{db: db}
}
