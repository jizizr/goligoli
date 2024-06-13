package config

import (
	"github.com/jizizr/goligoli/server/kitex_gen/live/liveservice"
	"github.com/jizizr/goligoli/server/kitex_gen/lottery/lotteryservice"
)

var (
	GlobalEtcdConfig   EtcdConfig
	GlobalServerConfig Config
)

var (
	LotteryClient lotteryservice.Client
	LiveClient    liveservice.Client
)
