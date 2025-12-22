package env

import (
	"os"
)

type Config struct {
	Debug bool
}

func LoadConfig() *Config {
	return &Config{
		Debug: os.Getenv("DEBUG") == "true",
	}
}
