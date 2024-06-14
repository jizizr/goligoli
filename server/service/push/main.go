package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	push "github.com/jizizr/goligoli/server/kitex_gen/push/pushservice"
	"github.com/jizizr/goligoli/server/service/push/config"
	"github.com/jizizr/goligoli/server/service/push/initialize"
	"github.com/jizizr/goligoli/server/service/push/initialize/rpc"
	"github.com/jizizr/goligoli/server/service/push/mq"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"log"
)

func main() {
	initialize.InitConfig()
	r, info := initialize.InitRegistry()
	publisher := initialize.InitProducer()
	subscriber := initialize.InitComsumer()
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint("localhost:4318"),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())
	rpc.Init()
	initialize.InitReciver()
	go func() {
		MessageSub := mq.NewSubscriberManager(subscriber)
		err := MessageSub.SubscribeMessageFromNsq(context.Background())
		if err != nil {
			klog.Error(err)
		}
	}()
	svr := push.NewServer(
		&PushServiceImpl{
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
