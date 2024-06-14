package main

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/jizizr/goligoli/server/common/consts"
	"github.com/jizizr/goligoli/server/service/record/config"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/client/v3"
	"log"
	"net"
	"time"
)

func main() {
	v := viper.New()
	v.SetConfigFile("../config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic("viper read config failed, err: " + err.Error())
	}
	etcdConf := new(config.EtcdConfig)
	err = v.Unmarshal(&etcdConf)
	if err != nil {
		panic("viper unmarshal config failed, err: " + err.Error())
	}
	c := config.Config{
		Name: consts.RecordSrv,
		Server: config.ServerConfig{
			Host: "127.0.0.1",
			Port: consts.RecordServerPort,
		},
		MinioInfo: config.MinioConfig{
			Endpoint: "127.0.0.1:9000",
			ID:       "RvT568oblWFE2dX5rj6K",
			Secret:   "4WKBSMcVy7KILrBo0iNqGma5CZLgqensFpYZMRDd",
			Bucket:   "video",
		},
		Stream: config.StreamConfig{
			Address: "http://127.0.0.1:7001/live",
		},
	}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{net.JoinHostPort(etcdConf.Host, etcdConf.Port)},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	configBytes, err := sonic.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	_, err = cli.Put(cli.Ctx(), etcdConf.Key, string(configBytes))

	resp, err := cli.Get(cli.Ctx(), etcdConf.Key)
	if err != nil {
		log.Fatal(err)
	}
	for _, kv := range resp.Kvs {
		fmt.Printf("%s : %s\n", kv.Key, kv.Value)
	}
}
