package mq

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/goligoli/server/kitex_gen/lottery"
	"github.com/jizizr/goligoli/server/service/delay/config"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

type PublisherManager struct {
	Publisher *nsq.Producer
}

func (p *PublisherManager) PushToNsq(id int64) error {
	body := []byte(strconv.FormatInt(id, 10))
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

func HandleMessage(message *nsq.Message) error {
	id, _ := strconv.ParseInt(string(message.Body), 10, 64)
	_, err := config.LotteryClient.DrawLottery(context.Background(), &lottery.DrawLotteryRequest{Id: id})
	if err != nil {
		klog.Errorf("draw lottery failed,%v", err)
	}
	return err
}

func NewPublisherManager(publisher *nsq.Producer) *PublisherManager {
	return &PublisherManager{Publisher: publisher}
}

func NewSubscriberManager(subscriber *nsq.Consumer) *SubscriberManager {
	return &SubscriberManager{Subscriber: subscriber}
}
