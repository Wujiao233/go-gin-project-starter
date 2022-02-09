package config

type Configuration struct {
	Server ServerConfig
}

type ServerConfig struct {
	Host       string
	ServerPort string `json:"server_port"`
}
