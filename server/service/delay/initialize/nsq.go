package initialize

import (
	"fmt"
	"github.com/jizizr/goligoli/server/service/delay/config"
	"github.com/nsqio/go-nsq"
)

func InitProducer() *nsq.Producer {
	nsqInfo := config.GlobalServerConfig.NsqInfo
	producer, err := nsq.NewProducer(fmt.Sprintf("%s:%d", nsqInfo.Host, nsqInfo.Port), nsq.NewConfig())
	if err != nil {
		panic("init nsq producer failed, err: " + err.Error())
	}
	return producer
}

func InitConsumer() *nsq.Consumer {
	nsqInfo := config.GlobalServerConfig.NsqInfo
	consumer, err := nsq.NewConsumer(nsqInfo.Topic, "1", nsq.NewConfig())
	if err != nil {
		panic("init nsq consumer failed, err: " + err.Error())
	}
	return consumer
}
