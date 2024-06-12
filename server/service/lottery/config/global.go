package config

import (
	"github.com/jizizr/goligoli/server/kitex_gen/delay/delaytaskservice"
	"github.com/jizizr/goligoli/server/kitex_gen/push/pushservice"
	"github.com/jizizr/goligoli/server/service/lottery/dao"
)

var (
	GlobalEtcdConfig   EtcdConfig
	GlobalServerConfig Config
)

var (
	WinnerDB *dao.Winner
)

var (
	DelayClient delaytaskservice.Client
	PushClient  pushservice.Client
)
