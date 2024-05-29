package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	user "github.com/jizizr/goligoli/server/kitex_gen/user/userservice"
	"github.com/jizizr/goligoli/server/service/user/config"
	"github.com/jizizr/goligoli/server/service/user/dao"
	"github.com/jizizr/goligoli/server/service/user/initialize"
	"log"
)

func main() {
	initialize.InitConfig()
	db := initialize.InitDB()
	r, info := initialize.InitRegistry()
	svr := user.NewServer(&UserServiceImpl{
		dao.NewUser(db),
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
