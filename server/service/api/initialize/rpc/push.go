package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	push "github.com/jizizr/goligoli/server/kitex_gen/push/pushservice"
	"github.com/jizizr/goligoli/server/service/api/biz/global"
	"github.com/jizizr/goligoli/server/service/api/config"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func initPush() {
	r, err := etcd.NewEtcdResolver([]string{net.JoinHostPort(config.GlobalEtcdConfig.Host, config.GlobalEtcdConfig.Port)})
	if err != nil {
		log.Fatal(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServiceConfig.Name),
		provider.WithExportEndpoint("localhost:4318"),
		provider.WithInsecure(),
	)
	c, err := push.NewClient(
		config.GlobalServiceConfig.PushSrv.Name,
		client.WithResolver(r),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: config.GlobalServiceConfig.PushSrv.Name,
			}),
	)
	if err != nil {
		log.Fatal(err)
	}
	global.PushClient = c
}
