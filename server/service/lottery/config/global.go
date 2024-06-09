package config

import (
	"github.com/jizizr/goligoli/server/service/lottery/dao"
)

var (
	GlobalEtcdConfig   EtcdConfig
	GlobalServerConfig Config
)

var (
	WinnerDB *dao.Winner
)
