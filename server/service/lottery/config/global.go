package config

import (
	"github.com/jizizr/goligoli/server/kitex_gen/delay/delaytaskservice"
	"github.com/jizizr/goligoli/server/kitex_gen/live/liveservice"
	"github.com/jizizr/goligoli/server/kitex_gen/push/pushservice"
	"github.com/jizizr/goligoli/server/service/lottery/dao"
)

var (
	GlobalEtcdConfig   EtcdConfig
	GlobalServerConfig Config
)

var (
	LotteryDB *dao.Lottery
)

var (
	DelayClient delaytaskservice.Client
	PushClient  pushservice.Client
	LiveClient  liveservice.Client
)
