package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	bullet "github.com/jizizr/goligoli/server/kitex_gen/bullet/bulletservice"
	"github.com/jizizr/goligoli/server/service/bullet/config"
	"github.com/jizizr/goligoli/server/service/bullet/dao"
	"github.com/jizizr/goligoli/server/service/bullet/initialize"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"log"
)

func main() {
	initialize.InitConfig()
	db := initialize.InitDB()
	r, info := initialize.InitRegistry()
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint("localhost:4318"),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())
	svr := bullet.NewServer(&BulletServiceImpl{
		dao.NewBullet(db),
	},
		server.WithServiceAddr(info.Addr),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: config.GlobalServerConfig.Name,
		}),
	)
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
