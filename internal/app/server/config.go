package server

type Config struct {
	Port        string `toml:"port"`
	LogLeveL    string `toml:"log_level"`
	DatabaseURL string `toml:database_url`
}

func NewConfig() *Config {
	return &Config{
		Port:     ":8000",
		LogLeveL: "debug",
	}
}
