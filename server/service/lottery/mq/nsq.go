package mq

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/goligoli/server/service/lottery/config"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"syscall"
)

type PublisherManager struct {
	Publisher *nsq.Producer
}

func (p *PublisherManager) PushToNsq(result []int64) error {
	body, err := sonic.Marshal(result)
	if err != nil {
		return err
	}
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

func HandleWinnersFunc(message *nsq.Message) error {
	var result []int64
	err := sonic.Unmarshal(message.Body, &result)
	if err != nil {
		klog.Errorf("unmarshal winners failed, %v", err)
		return err
	}
	err = config.WinnerDB.AddWinners(result[0], result[1:])
	if err != nil {
		klog.Errorf("add winners failed, %v", err)
	}
	return err
}

func NewPublisherManager(publisher *nsq.Producer) *PublisherManager {
	return &PublisherManager{Publisher: publisher}
}

func NewSubscriberManager(subscriber *nsq.Consumer) *SubscriberManager {
	return &SubscriberManager{Subscriber: subscriber}
}
