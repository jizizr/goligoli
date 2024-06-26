package mq

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/goligoli/server/kitex_gen/base"
	message2 "github.com/jizizr/goligoli/server/kitex_gen/message"
	"github.com/jizizr/goligoli/server/kitex_gen/push"
	"github.com/jizizr/goligoli/server/service/push/config"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"syscall"
)

type PublisherManager struct {
	Publisher *nsq.Producer
}

func (p PublisherManager) PushMessageToNsq(ctx context.Context, request *push.PushMessageRequest) error {
	body, err := sonic.Marshal(request.Message)
	if err != nil {
		return err
	}
	return p.Publisher.Publish(config.GlobalServerConfig.NsqInfo.Topic, body)
}

type SubscriberManager struct {
	Subscriber *nsq.Consumer
}

func (s SubscriberManager) SubscribeMessageFromNsq(ctx context.Context) error {
	nsqInfo := config.GlobalServerConfig.NsqInfo
	s.Subscriber.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		var req base.LiveMessage
		err := sonic.Unmarshal(message.Body, &req)
		if err != nil {
			klog.Errorf("subscribe unmarshal message failed, %v", err)
			return err
		}
		err = config.MessageClient.AddMessage(ctx, &message2.AddMessageRequest{
			Message: &req,
		})
		if err != nil {
			klog.Errorf("subscribe add message failed, %v", err)
			return err
		}
		return nil
	}))
	err := s.Subscriber.ConnectToNSQD(fmt.Sprintf("%s:%d", nsqInfo.Host, nsqInfo.Port))
	if err != nil {
		klog.Errorf("subscribe connect to nsqd failed, %v", err)
		return err
	}

	// 处理退出信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// 停止消费者连接
	s.Subscriber.Stop()
	<-s.Subscriber.StopChan

	return nil
}

func NewPublisherManager(publisher *nsq.Producer) *PublisherManager {
	return &PublisherManager{Publisher: publisher}
}

func NewSubscriberManager(subscriber *nsq.Consumer) *SubscriberManager {
	return &SubscriberManager{Subscriber: subscriber}
}
