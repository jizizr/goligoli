package initialize

import (
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/goligoli/server/common/consts"
	"github.com/jizizr/goligoli/server/service/message/config"
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"net"
	"time"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile(consts.ConfigPath)
	err := v.ReadInConfig()
	if err != nil {
		panic("viper read config failed, err: " + err.Error())
	}
	err = v.Unmarshal(&config.GlobalEtcdConfig)
	if err != nil {
		panic("viper unmarshal config failed, err: " + err.Error())
	}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{net.JoinHostPort(config.GlobalEtcdConfig.Host, config.GlobalEtcdConfig.Port)},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		klog.Fatalf("new etcd client failed, err: %s", err)
	}
	defer cli.Close()
	content, err := cli.Get(cli.Ctx(), config.GlobalEtcdConfig.Key)
	if err != nil {
		log.Fatalf("etcd get config failed, err: %s", err)
	}
	kv := content.Kvs[0]
	err = sonic.Unmarshal(kv.Value, &config.GlobalServerConfig)
	if err != nil {
		klog.Fatalf("sonic unmarshal config failed, err: %s", err)
	}
	klog.Infof("get config from etcd success,key: %s, config: %v", kv.Key, config.GlobalServerConfig)
}
