package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	Message "github.com/jizizr/goligoli/server/kitex_gen/message/messageservice"
	"github.com/jizizr/goligoli/server/service/push/config"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func initMessage() {
	r, err := etcd.NewEtcdResolver([]string{net.JoinHostPort(config.GlobalEtcdConfig.Host, config.GlobalEtcdConfig.Port)})
	if err != nil {
		log.Fatal(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint("localhost:4318"),
		provider.WithInsecure(),
	)
	c, err := Message.NewClient(
		config.GlobalServerConfig.MessageSrv.Name,
		client.WithSuite(tracing.NewClientSuite()),
		client.WithResolver(r),
		client.WithClientBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: config.GlobalServerConfig.MessageSrv.Name,
			}),
	)
	if err != nil {
		log.Fatal(err)
	}
	config.MessageClient = c
}
