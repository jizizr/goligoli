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
)

const (
	ErrPassword = "ErrorPassword"
)

const (
	MySqlAlreadyExist = "MySqlAlreadyExist"
	MySqlNotExist
	MySqlInsertError
)
