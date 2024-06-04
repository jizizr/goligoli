package config

import (
	"github.com/jizizr/goligoli/server/kitex_gen/message/messageservice"
	"sync"
)

var (
	GlobalEtcdConfig   EtcdConfig
	GlobalServerConfig Config
)

var (
	MessageClient messageservice.Client
	Receiver      sync.Map
)
