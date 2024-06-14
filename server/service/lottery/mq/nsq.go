package mq

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/goligoli/server/kitex_gen/base"
	"github.com/jizizr/goligoli/server/service/lottery/config"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"syscall"
)

type PublisherManager struct {
	Publisher *nsq.Producer
}

func (p *PublisherManager) PushToNsq(body []byte) error {
	return p.Publisher.Publish(config.GlobalServerConfig.NsqInfo.Topic, body)
}

type SubscriberManager struct {
	Subscriber *nsq.Consumer
}

func (s *SubscriberManager) SubscribeFromNsq(f nsq.HandlerFunc) error {
	nsqInfo := config.GlobalServerConfig.NsqInfo
	s.Subscriber.AddHandler(f)
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

func HandleFunc(message *nsq.Message) error {
	var gift base.Gift
	err := sonic.Unmarshal(message.Body, &gift)
	if err != nil {
		klog.Errorf("unmarshal winners failed, %v", err)
		return err
	}
	err = config.LotteryDB.SetLotteryDB(&gift)
	if err != nil {
		klog.Errorf("set lottery db failed, %v", err)
	}
	return err
}

func NewPublisherManager(publisher *nsq.Producer) *PublisherManager {
	return &PublisherManager{Publisher: publisher}
}

func NewSubscriberManager(subscriber *nsq.Consumer) *SubscriberManager {
	return &SubscriberManager{Subscriber: subscriber}
}
