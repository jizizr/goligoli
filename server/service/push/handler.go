package main

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/goligoli/server/common/tools"
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
	PushMessageToNsq(ctx context.Context, request *push.PushMessageRequest) error
}

// PushMessage implements the PushServiceImpl interface.
func (s *PushServiceImpl) PushMessage(ctx context.Context, req *push.PushMessageRequest) (err error) {
	r, ok := config.Receiver.Load(req.Message.LiveId)
	if !ok {
		return errors.New("live not exist")
	}
	r.(*sync.Map).Range(func(key, value interface{}) bool {
		rec := value.(chan *base.LiveMessage)
		rec <- req.Message
		return true
	})
	err = s.PushMessageToNsq(ctx, &push.PushMessageRequest{Message: req.Message})
	if err != nil {
		klog.Errorf("push message to nsq failed, %v", err)
	}
	return
}

func (s *PushServiceImpl) ReceiveMessage(stream push.PushService_ReceiveMessageServer) (err error) {
	rec := make(chan *base.LiveMessage)
	req, err := stream.Recv()
	defer stream.Close()
	if !tools.CheckLiveRoom(req.LiveId, &config.LiveClient) {
		return errors.New("live room not exist")
	}
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
		case Message, ok := <-rec:
			if !ok {
				goto EXIT
			}
			err = stream.Send(&push.ReceiveMessageResponse{Message: Message})
			if err != nil {
				goto EXIT
			}
		}
	}
EXIT:
	v.(*sync.Map).Delete(req.UserId)
	return
}
