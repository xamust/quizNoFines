package service

type Config struct {
	LogLevel string `toml:"log_level"`
	PortgRPC string `toml:"port_grpc"`
	Port     string `toml:"port"`
}

func NewConfig() *Config {
	return &Config{
		LogLevel: "info", //default param
	}
}
