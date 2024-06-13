package config

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"dao"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type NsqConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Topic    string `json:"topic"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type RPCSrvConfig struct {
	Name string `json:"name"`
}

type Config struct {
	Name      string       `json:"name"`
	Server    ServerConfig `json:"server"`
	MysqlInfo MysqlConfig  `json:"mysql"`
	RedisInfo RedisConfig  `json:"redis"`
	NsqInfo   NsqConfig    `json:"nsq"`
	DelaySrv  RPCSrvConfig `json:"delay_srv"`
	PushSrv   RPCSrvConfig `json:"push_srv"`
	LiveSrv   RPCSrvConfig `json:"live_srv"`
}

type EtcdConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port string `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}
