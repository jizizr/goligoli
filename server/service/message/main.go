package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	Message "github.com/jizizr/goligoli/server/kitex_gen/message/messageservice"
	"github.com/jizizr/goligoli/server/service/message/config"
	"github.com/jizizr/goligoli/server/service/message/dao"
	"github.com/jizizr/goligoli/server/service/message/initialize"
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
	svr := Message.NewServer(
		&MessageServiceImpl{
			dao.NewMessage(db),
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
