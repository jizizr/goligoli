package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/jizizr/goligoli/server/common/consts"
	user "github.com/jizizr/goligoli/server/kitex_gen/user/userservice"
	"github.com/jizizr/goligoli/server/service/user/dao"
	"github.com/jizizr/goligoli/server/service/user/initialize"
	"log"
	"net"
	"strconv"
)

func main() {
	initialize.InitConfig()
	db := initialize.InitDB()
	r, info := initialize.InitRegistry()
	svr := user.NewServer(&UserServiceImpl{
		dao.NewUser(db),
	},
		server.WithServiceAddr(utils.NewNetAddr("tcp", net.JoinHostPort("127.0.0.1", strconv.Itoa(consts.UserServerPort)))),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "user_srv",
		}),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
