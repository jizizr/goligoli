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

type RecChan struct {
	rec chan *base.LiveMessage
	fl  chan struct{}
}

func NewRecChan() *RecChan {
	return &RecChan{
		rec: make(chan *base.LiveMessage),
		fl:  make(chan struct{}),
	}
}

func (r *RecChan) Get() (msg chan *base.LiveMessage, fl chan struct{}) {
	return r.rec, r.fl
}

// PushMessage implements the PushServiceImpl interface.
func (s *PushServiceImpl) PushMessage(ctx context.Context, req *push.PushMessageRequest) (err error) {
	r, ok := config.Receiver.Load(req.Message.LiveId)
	if !ok {
		return errors.New("live not exist")
	}
	if config.Limiter.ShouldSendWord(req.Message.Content) {
		r.(*sync.Map).Range(func(key, value interface{}) bool {
			rec, flag := value.(*RecChan).Get()
			select {
			case rec <- req.Message:
			case <-flag:
			}
			return true
		})
	}
	err = s.PushMessageToNsq(ctx, &push.PushMessageRequest{Message: req.Message})
	if err != nil {
		klog.Errorf("push message to nsq failed, %v", err)
	}
	return
}

func (s *PushServiceImpl) ReceiveMessage(stream push.PushService_ReceiveMessageServer) (err error) {
	req, err := stream.Recv()
	defer stream.Close()
	if !tools.CheckLiveRoom(req.LiveId, &config.LiveClient) {
		return errors.New("live room not exist")
	}
	v, okk := config.Receiver.Load(req.LiveId)
	if !okk {
		return errors.New("live not exist")
	}
	recChan := NewRecChan()
	rec, flag := recChan.Get()
	v.(*sync.Map).Store(req.UserId, recChan)
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
	close(flag)
	return
}

// StopMessage implements the PushServiceImpl interface.
func (s *PushServiceImpl) StopMessage(ctx context.Context, req *push.StopMessageRequest) (err error) {
	r, ok := config.Receiver.Load(req.LiveId)
	if !ok {
		return errors.New("live not exist")
	}
	config.Receiver.Delete(req.LiveId)
	go r.(*sync.Map).Range(func(key, value interface{}) bool {
		_, flag := value.(*RecChan).Get()
		close(flag)
		return true
	})
	return
}

// InitLiveRoomReciver implements the PushServiceImpl interface.
func (s *PushServiceImpl) InitLiveRoomReciver(ctx context.Context, req *push.InitLiveRoomReciverRequest) (err error) {
	// TODO: Your code here...
	return
}
