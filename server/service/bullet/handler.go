package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/goligoli/server/kitex_gen/base"
	bullet "github.com/jizizr/goligoli/server/kitex_gen/bullet"
)

// BulletServiceImpl implements the last service interface defined in the IDL.
type BulletServiceImpl struct {
	MySqlServiceImpl
}

type MySqlServiceImpl interface {
	CreateBullet(bullet *base.Bullet) error
	GetBulletByID(id int64) (*base.Bullet, error)
	GetHistoryBulletsByTime(liveID int64, startTime int64, offset int64) ([]*base.Bullet, error)
}

// AddBullet implements the BulletServiceImpl interface.
func (s *BulletServiceImpl) AddBullet(ctx context.Context, req *bullet.AddBulletRequest) (err error) {
	err = s.CreateBullet(req.Bullet)
	return
}

// GetBullet implements the BulletServiceImpl interface.
func (s *BulletServiceImpl) GetBullet(ctx context.Context, req *bullet.GetBulletRequest) (resp *bullet.GetBulletResponse, err error) {
	resp = new(bullet.GetBulletResponse)
	bul, err := s.GetBulletByID(req.BulletId)
	if err != nil {
		klog.Errorf("get bullet by id failed, %v", err)
		return
	}
	resp.Bullet = bul
	return
}

// GetHistoryBullets implements the BulletServiceImpl interface.
func (s *BulletServiceImpl) GetHistoryBullets(ctx context.Context, req *bullet.GetHistoryBulletsRequest) (resp *bullet.GetHistoryBulletsResponse, err error) {
	resp = new(bullet.GetHistoryBulletsResponse)
	buls, err := s.GetHistoryBulletsByTime(req.LiveId, req.StartTime, req.Offset)
	if err != nil {
		klog.Errorf("get history bullets by time failed, %v", err)
		return
	}
	resp.Bullets = buls
	return
}
