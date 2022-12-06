package server

type Config struct {
	Port     string `toml:"port"`
	LogLeveL string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		Port:     ":8000",
		LogLeveL: "debug",
	}
}
