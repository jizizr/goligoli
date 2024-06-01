package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	bullet "github.com/jizizr/goligoli/server/kitex_gen/bullet/bulletservice"
	"github.com/jizizr/goligoli/server/service/api/biz/global"
	"github.com/jizizr/goligoli/server/service/api/config"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func initBullet() {
	r, err := etcd.NewEtcdResolver([]string{net.JoinHostPort(config.GlobalEtcdConfig.Host, config.GlobalEtcdConfig.Port)})
	if err != nil {
		log.Fatal(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServiceConfig.Name),
		provider.WithExportEndpoint("localhost:4318"),
		provider.WithInsecure(),
	)
	c, err := bullet.NewClient(
		config.GlobalServiceConfig.BulletSrv.Name,
		client.WithResolver(r),
		client.WithClientBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: config.GlobalServiceConfig.BulletSrv.Name,
			}),
		client.WithSuite(tracing.NewClientSuite()),
	)
	if err != nil {
		log.Fatal(err)
	}
	global.BulletClient = c
}
