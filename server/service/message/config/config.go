package config

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"dao"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Config struct {
	Name      string       `json:"name"`
	Server    ServerConfig `json:"server"`
	MysqlInfo MysqlConfig  `json:"mysql"`
	LiveSrv   RPCSrvConfig `json:"live_srv"`
}

type RPCSrvConfig struct {
	Name string `json:"name"`
}

type EtcdConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port string `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}
