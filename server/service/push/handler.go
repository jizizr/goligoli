package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	push "github.com/jizizr/goligoli/server/kitex_gen/push"
)

// PushServiceImpl implements the last service interface defined in the IDL.
type PushServiceImpl struct {
	NsqServiceImpl
}

type NsqServiceImpl interface {
	PushBulletToNsq(ctx context.Context, request *push.PushBulletRequest) error
}

// PushBullet implements the PushServiceImpl interface.
func (s *PushServiceImpl) PushBullet(ctx context.Context, req *push.PushBulletRequest) (err error) {
	err = s.PushBulletToNsq(ctx, &push.PushBulletRequest{Bullet: req.Bullet})
	if err != nil {
		klog.Errorf("push bullet to nsq failed, %v", err)
	}
	return
}
