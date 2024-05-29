package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	bullet "github.com/jizizr/goligoli/server/kitex_gen/bullet/bulletservice"
	"github.com/jizizr/goligoli/server/service/bullet/config"
	"github.com/jizizr/goligoli/server/service/bullet/initialize"
	"log"
)

func main() {
	initialize.InitConfig()
	r, info := initialize.InitRegistry()
	svr := bullet.NewServer(new(BulletServiceImpl),
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
