package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	lottery "github.com/jizizr/goligoli/server/kitex_gen/lottery/lotteryservice"
	"github.com/jizizr/goligoli/server/service/delay/config"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func initLottery() {
	r, err := etcd.NewEtcdResolver([]string{net.JoinHostPort(config.GlobalEtcdConfig.Host, config.GlobalEtcdConfig.Port)})
	if err != nil {
		log.Fatal(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint("localhost:4318"),
		provider.WithInsecure(),
	)
	c, err := lottery.NewClient(
		config.GlobalServerConfig.LotterySrv.Name,
		client.WithResolver(r),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: config.GlobalServerConfig.LotterySrv.Name,
			}),
	)
	if err != nil {
		log.Fatal(err)
	}
	config.LotteryClient = c
}
