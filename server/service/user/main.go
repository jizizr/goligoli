package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	user "github.com/jizizr/goligoli/server/kitex_gen/user/userservice"
	"github.com/jizizr/goligoli/server/service/user/config"
	"github.com/jizizr/goligoli/server/service/user/dao"
	"github.com/jizizr/goligoli/server/service/user/initialize"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
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
	svr := user.NewServer(&UserServiceImpl{
		dao.NewUser(db),
	},
		server.WithServiceAddr(info.Addr),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: config.GlobalServerConfig.Name,
		}),
		server.WithSuite(tracing.NewServerSuite()),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
