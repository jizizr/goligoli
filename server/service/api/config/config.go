package config

type EtcdConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port string `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

type Config struct {
	Name      string       `json:"name"`
	UserSrv   RPCSrvConfig `json:"user_srv"`
	BulletSrv RPCSrvConfig `json:"bullet_srv"`
	PushSrv   RPCSrvConfig `json:"push_srv"`
}

type RPCSrvConfig struct {
	Name string `json:"name"`
}
