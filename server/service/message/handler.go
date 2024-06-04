package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/goligoli/server/kitex_gen/base"
	Message "github.com/jizizr/goligoli/server/kitex_gen/message"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct {
	MySqlServiceImpl
}

type MySqlServiceImpl interface {
	CreateMessage(Message *base.LiveMessage) error
	GetMessageByID(id int64) (*base.LiveMessage, error)
	GetHistoryMessagesByTime(liveID int64, startTime int64, offset int64) ([]*base.LiveMessage, error)
}

// AddMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) AddMessage(ctx context.Context, req *Message.AddMessageRequest) (err error) {
	err = s.CreateMessage(req.Message)
	return
}

// GetMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetMessage(ctx context.Context, req *Message.GetMessageRequest) (resp *Message.GetMessageResponse, err error) {
	resp = new(Message.GetMessageResponse)
	bul, err := s.GetMessageByID(req.Id)
	if err != nil {
		klog.Errorf("get message by id failed, %v", err)
		return
	}
	resp.Message = bul
	return
}

// GetHistoryMessages implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetHistoryMessages(ctx context.Context, req *Message.GetHistoryMessagesRequest) (resp *Message.GetHistoryMessagesResponse, err error) {
	resp = new(Message.GetHistoryMessagesResponse)
	buls, err := s.GetHistoryMessagesByTime(req.LiveId, req.StartTime, req.Offset)
	if err != nil {
		klog.Errorf("get history Messages by time failed, %v", err)
		return
	}
	resp.Messages = buls
	return
}
