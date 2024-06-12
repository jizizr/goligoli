package initialize

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/jizizr/goligoli/server/common/consts"
	"github.com/jizizr/goligoli/server/service/lottery/config"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
	"strconv"
)

// InitRegistry to init etcd
func InitRegistry() (registry.Registry, *registry.Info) {
	// 使用 etcd 注册
	r, err := etcd.NewEtcdRegistry(
		[]string{net.JoinHostPort(config.GlobalEtcdConfig.Host, config.GlobalEtcdConfig.Port)},
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
		ServiceName: config.GlobalServerConfig.Name,
		Addr:        utils.NewNetAddr("tcp", net.JoinHostPort("127.0.0.1", strconv.Itoa(config.GlobalServerConfig.Server.Port))),
		Tags: map[string]string{
			"ID": sf.Generate().Base36(),
		},
	}
	return r, info
}
