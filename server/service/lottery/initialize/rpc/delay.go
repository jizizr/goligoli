package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	delay "github.com/jizizr/goligoli/server/kitex_gen/delay/delaytaskservice"
	"github.com/jizizr/goligoli/server/service/lottery/config"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func initDelay() {
	r, err := etcd.NewEtcdResolver([]string{net.JoinHostPort(config.GlobalEtcdConfig.Host, config.GlobalEtcdConfig.Port)})
	if err != nil {
		log.Fatal(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint("localhost:4318"),
		provider.WithInsecure(),
	)
	c, err := delay.NewClient(
		config.GlobalServerConfig.DelaySrv.Name,
		client.WithResolver(r),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: config.GlobalServerConfig.DelaySrv.Name,
			}),
	)
	if err != nil {
		log.Fatal(err)
	}
	config.DelayClient = c
}
