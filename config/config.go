package config

type Config struct {
	HttpHost    string
	HttpPort    int
	DatabaseUrl string
}

func GenerateDefaultConfig() *Config {
	return &Config{HttpHost: "0.0.0.0", HttpPort: 8080, DatabaseUrl: ""}
}
