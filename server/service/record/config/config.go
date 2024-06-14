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

type MinioConfig struct {
	Endpoint string `json:"endpoint"`
	ID       string `json:"id"`
	Secret   string `json:"secret"`
	Bucket   string `json:"bucket"`
}

type StreamConfig struct {
	Address string `json:"address"`
}

type Config struct {
	Name      string       `json:"name"`
	Server    ServerConfig `json:"server"`
	MinioInfo MinioConfig  `json:"minio"`
	Stream    StreamConfig `json:"stream"`
}
