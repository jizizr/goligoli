package consts

const (
	ConfigPath = "config.yaml"

	MySqlDSN = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

const (
	UserSnowflakeNode = iota + 1
	MessageSnowflakeNode
	EtcdSnowflakeNode
)

const (
	ApiServerPort = 8000 + iota
	UserServerPort
	MessageServerPort
	PushServerPort
	LiveServerPort
)

const (
	ApiSrv     = "api_srv"
	UserSrv    = "user_srv"
	MessageSrv = "message_srv"
	PushSrv    = "push_srv"
	LiveSrv    = "live_srv"
)

const (
	BULLET = iota
	GIFT
	SUPERMSG
)
