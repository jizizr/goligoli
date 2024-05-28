package initialize

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/jizizr/goligoli/server/common/consts"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
	"strconv"
)

// InitRegistry to init etcd
func InitRegistry() (registry.Registry, *registry.Info) {
	// 使用 etcd 注册
	r, err := etcd.NewEtcdRegistry(
		[]string{"127.0.0.1:2379"},
	)
	if err != nil {
		klog.Fatalf("new etcd register failed: %s", err.Error())
	}

	// 使用 snowflake 生成服务名称
	sf, err := snowflake.NewNode(consts.EtcdSnowflakeNode)
	if err != nil {
		klog.Fatalf("generate service name failed: %s", err.Error())
	}
	info := &registry.Info{
		ServiceName: "user_srv",
		Addr:        utils.NewNetAddr("tcp", net.JoinHostPort("127.0.0.1", strconv.Itoa(consts.UserServerPort))),
		Tags: map[string]string{
			"ID": sf.Generate().Base36(),
		},
	}
	return r, info
}
