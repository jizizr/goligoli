package main

import (
	"context"
	bullet "github.com/jizizr/goligoli/server/kitex_gen/bullet"
)

// BulletServiceImpl implements the last service interface defined in the IDL.
type BulletServiceImpl struct{}

// AddBullet implements the BulletServiceImpl interface.
func (s *BulletServiceImpl) AddBullet(ctx context.Context, request *bullet.AddBulletRequest) (resp *bullet.AddBulletResponse, err error) {
	// TODO: Your code here...
	return
}

// GetBullet implements the BulletServiceImpl interface.
func (s *BulletServiceImpl) GetBullet(ctx context.Context, request *bullet.GetBulletRequest) (resp *bullet.GetBulletResponse, err error) {
	// TODO: Your code here...
	return
}
