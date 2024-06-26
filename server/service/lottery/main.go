package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	lottery "github.com/jizizr/goligoli/server/kitex_gen/lottery/lotteryservice"
	"github.com/jizizr/goligoli/server/service/lottery/config"
	"github.com/jizizr/goligoli/server/service/lottery/dao"
	"github.com/jizizr/goligoli/server/service/lottery/initialize"
	"github.com/jizizr/goligoli/server/service/lottery/initialize/rpc"
	"github.com/jizizr/goligoli/server/service/lottery/mq"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"log"
)

func main() {
	initialize.InitConfig()
	db := initialize.InitMySql()
	rd := initialize.InitRedis()
	publisher := initialize.InitProducer()
	subscriber := initialize.InitConsumer()
	r, info := initialize.InitRegistry()
	rpc.Init()
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint("localhost:4318"),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())
	config.LotteryDB = dao.NewLottery(db)
	go func() {
		MessageSub := mq.NewSubscriberManager(subscriber)
		err := MessageSub.SubscribeFromNsq(mq.HandleFunc)
		if err != nil {
			klog.Error(err)
		}
	}()
	svr := lottery.NewServer(
		&LotteryServiceImpl{
			dao.NewWinner(db),
			config.LotteryDB,
			dao.NewLotteryRedis(rd),
			mq.NewPublisherManager(publisher),
		},
		server.WithServiceAddr(info.Addr),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: config.GlobalServerConfig.Name,
		}),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
