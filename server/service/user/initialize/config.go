package initialize

import (
	"github.com/jizizr/goligoli/server/common/consts"
	"github.com/jizizr/goligoli/server/service/user/config"
	"github.com/spf13/viper"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile(consts.ConfigPath)
	err := v.ReadInConfig()
	if err != nil {
		panic("viper read config failed, err: " + err.Error())
	}
	err = v.Unmarshal(&config.GlobalServerConfig)
	if err != nil {
		panic("viper unmarshal config failed, err: " + err.Error())
	}
}
