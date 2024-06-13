package main

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/jizizr/goligoli/server/common/consts"
	"github.com/jizizr/goligoli/server/service/live/config"
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
		Name: consts.LiveSrv,
		Server: config.ServerConfig{
			Host: "127.0.0.1",
			Port: consts.LiveServerPort,
		},
		MysqlInfo: config.MysqlConfig{
			Host:     "127.0.0.1",
			Port:     3306,
			Username: "root",
			Password: "",
			Name:     "goligoli",
		},
		RedisInfo: config.RedisConfig{
			Host:     "127.0.0.1",
			Port:     6379,
			Password: "",
			DB:       0,
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
