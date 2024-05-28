package rpc

import (
	"github.com/cloudwego/kitex/client"
	user "github.com/jizizr/goligoli/server/kitex_gen/user/userservice"
	"github.com/jizizr/goligoli/server/service/api/biz/global"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

// initUser to init user service
func initUser() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	c, err := user.NewClient("user_srv", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	global.UserClient = c
}
