package config

import (
	"github.com/jizizr/goligoli/server/common/tools"
	"github.com/jizizr/goligoli/server/kitex_gen/live/liveservice"
	"github.com/jizizr/goligoli/server/kitex_gen/message/messageservice"
	"sync"
	"time"
)

var (
	GlobalEtcdConfig   EtcdConfig
	GlobalServerConfig Config
)

var (
	MessageClient messageservice.Client
	LiveClient    liveservice.Client
	Receiver      sync.Map
	Limiter       = tools.NewLimiter(time.Second, 20)
)
