package config

type EtcdConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port string `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

type Config struct {
	Name       string       `json:"name"`
	UserSrv    RPCSrvConfig `json:"user_srv"`
	MessageSrv RPCSrvConfig `json:"message_srv"`
	PushSrv    RPCSrvConfig `json:"push_srv"`
	LiveSrv    RPCSrvConfig `json:"live_srv"`
}

type RPCSrvConfig struct {
	Name string `json:"name"`
}
