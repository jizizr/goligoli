package main

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/goligoli/server/kitex_gen/base"
	"github.com/jizizr/goligoli/server/kitex_gen/push"
	"github.com/jizizr/goligoli/server/service/push/config"
	"sync"
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
	r, ok := config.Receiver.Load(req.Bullet.LiveId)
	if !ok {
		return errors.New("live not exist")
	}
	r.(*sync.Map).Range(func(key, value interface{}) bool {
		rec := value.(chan *base.Bullet)
		rec <- req.Bullet
		return true
	})
	err = s.PushBulletToNsq(ctx, &push.PushBulletRequest{Bullet: req.Bullet})
	if err != nil {
		klog.Errorf("push bullet to nsq failed, %v", err)
	}
	return
}

func (s *PushServiceImpl) ReceiveBullet(stream push.PushService_ReceiveBulletServer) (err error) {
	rec := make(chan *base.Bullet)
	req, err := stream.Recv()
	v, _ := config.Receiver.LoadOrStore(req.LiveId, &sync.Map{})
	v.(*sync.Map).Store(req.UserId, rec)
	ok := make(chan struct{})
	go func() {
		stream.Recv()
		close(ok)
	}()

	for {
		select {
		case <-ok:
			goto EXIT
		case <-stream.Context().Done():
			goto EXIT
		case bullet := <-rec:
			err = stream.Send(&push.ReceiveBulletResponse{Bullet: bullet})
			if err != nil {
				goto EXIT
			}
		}
	}
EXIT:
	v.(*sync.Map).Delete(req.UserId)
	return
}
