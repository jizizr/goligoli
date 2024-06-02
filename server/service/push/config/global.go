package config

import "github.com/jizizr/goligoli/server/kitex_gen/bullet/bulletservice"

var (
	GlobalEtcdConfig   EtcdConfig
	GlobalServerConfig Config
)

var (
	BulletClient bulletservice.Client
)
