package config

import "github.com/jizizr/goligoli/server/kitex_gen/live/liveservice"

var (
	GlobalServerConfig Config
	GlobalEtcdConfig   EtcdConfig
)

var (
	LiveClient *liveservice.Client
)
