package config

import (
	"github.com/jizizr/goligoli/server/kitex_gen/bullet/bulletservice"
	"sync"
)

var (
	GlobalEtcdConfig   EtcdConfig
	GlobalServerConfig Config
)

var (
	BulletClient bulletservice.Client
	Receiver     sync.Map
)
