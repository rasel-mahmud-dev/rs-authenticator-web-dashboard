package config

import (
	"os"
	"sync"
)

type Config struct {
	Port string
}

var config *Config
var once sync.Once

func InitConfig() {
	once.Do(func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}

		config = &Config{
			Port: port,
		}
	})
}

func ConfigInstance() *Config {
	InitConfig()
	return config
}
