package store

type Config struct {
	DatabasURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}
