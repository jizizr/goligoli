package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	delay "github.com/jizizr/goligoli/server/kitex_gen/delay/delaytaskservice"
	"github.com/jizizr/goligoli/server/service/delay/config"
	"github.com/jizizr/goligoli/server/service/delay/dqueue"
	"github.com/jizizr/goligoli/server/service/delay/initialize"
	"github.com/jizizr/goligoli/server/service/delay/initialize/rpc"
	"github.com/jizizr/goligoli/server/service/delay/mq"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"log"
	"time"
)

func main() {
	initialize.InitConfig()
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
	go func() {
		MessageSub := mq.NewSubscriberManager(subscriber)
		err := MessageSub.SubscribeFromNsq(mq.HandleMessage)
		if err != nil {
			klog.Error(err)
		}
	}()
	pub := mq.NewPublisherManager(publisher)
	dq := dqueue.NewDelayQueue("lottery", time.Second, rd, pub.PushToNsq)
	dq.Start()
	defer dq.Stop()
	svr := delay.NewServer(
		&DelayTaskServiceImpl{
			dq,
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
