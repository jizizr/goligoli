package rpc

import (
	"github.com/cloudwego/kitex/client"
	bullet "github.com/jizizr/goligoli/server/kitex_gen/bullet/bulletservice"
	"github.com/jizizr/goligoli/server/service/api/biz/global"
	"github.com/jizizr/goligoli/server/service/api/config"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func initBullet() {
	r, err := etcd.NewEtcdResolver([]string{net.JoinHostPort(config.GlobalEtcdConfig.Host, config.GlobalEtcdConfig.Port)})
	if err != nil {
		log.Fatal(err)
	}
	c, err := bullet.NewClient(config.GlobalServiceConfig.BulletSrv.Name, client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	global.BulletClient = c
}
