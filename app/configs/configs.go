package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"sync"
)

type Config struct {
	Port              string
	DATABASE_HOST     string
	DATABASE_PORT     int32
	DATABASE_USER     string
	DATABASE_PASSWORD string
	DATABASE_NAME     string
	CACHE_STORAGE     string
}

var config *Config
var once sync.Once

func init() {
	//if os.Getenv("APP_ENV") == "development" {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
	//}

	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		log.Fatalf("Invalid DATABASE_PORT value: %v", err)
	}

	once.Do(func() {
		config = &Config{
			Port:              os.Getenv("PORT"),
			DATABASE_HOST:     os.Getenv("DATABASE_HOST"),
			DATABASE_PORT:     int32(port),
			DATABASE_USER:     os.Getenv("DATABASE_USER"),
			DATABASE_PASSWORD: os.Getenv("DATABASE_PASSWORD"),
			DATABASE_NAME:     os.Getenv("DATABASE_NAME"),
			CACHE_STORAGE:     os.Getenv("CACHE_STORAGE"),
		}
	})
}

func ConfigInstance() *Config {
	return config
}
