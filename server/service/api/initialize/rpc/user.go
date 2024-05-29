package rpc

import (
	"github.com/cloudwego/kitex/client"
	user "github.com/jizizr/goligoli/server/kitex_gen/user/userservice"
	"github.com/jizizr/goligoli/server/service/api/biz/global"
	"github.com/jizizr/goligoli/server/service/api/config"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func initUser() {
	r, err := etcd.NewEtcdResolver([]string{net.JoinHostPort(config.GlobalEtcdConfig.Host, config.GlobalEtcdConfig.Port)})
	if err != nil {
		log.Fatal(err)
	}
	c, err := user.NewClient(config.GlobalServiceConfig.UserSrv.Name, client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	global.UserClient = c
}
