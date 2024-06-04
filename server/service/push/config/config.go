package config

type EtcdConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port string `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type NsqConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Topic    string `json:"topic"`
}

type RPCSrvConfig struct {
	Name string `json:"name"`
}

type Config struct {
	Name       string       `json:"name"`
	Server     ServerConfig `json:"server"`
	NsqInfo    NsqConfig    `json:"nsq"`
	MessageSrv RPCSrvConfig `json:"message_srv"`
	ApiSrv     RPCSrvConfig `json:"api_srv"`
}
