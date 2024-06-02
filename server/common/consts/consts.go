package consts

const (
	ConfigPath = "config.yaml"

	MySqlDSN = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

const (
	UserSnowflakeNode = iota + 1
	BulletSnowflakeNode
	EtcdSnowflakeNode
)

const (
	ApiServerPort = 8000 + iota
	UserServerPort
	BulletServerPort
	PushServerPort
)

const (
	ApiSrv    = "api_srv"
	UserSrv   = "user_srv"
	BulletSrv = "bullet_srv"
	PushSrv   = "push_srv"
)
