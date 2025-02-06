package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type configI struct {
	Port              string
	DATABASE_HOST     string
	DATABASE_PORT     int32
	DATABASE_USER     string
	DATABASE_PASSWORD string
	DATABASE_NAME     string
	CACHE_STORAGE     string
	JWT_SECRET_KEY    string
	APP_LOGO_URL      string
}

var Config *configI

func init() {
	//os.Setenv("APP_ENV", "development")
	//if os.Getenv("APP_ENV") == "development" {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
	//}

	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		log.Fatalf("Invalid DATABASE_PORT value: %v", err)
	}

	Config = &configI{
		Port:              os.Getenv("PORT"),
		DATABASE_HOST:     os.Getenv("DATABASE_HOST"),
		DATABASE_PORT:     int32(port),
		DATABASE_USER:     os.Getenv("DATABASE_USER"),
		DATABASE_PASSWORD: os.Getenv("DATABASE_PASSWORD"),
		DATABASE_NAME:     os.Getenv("DATABASE_NAME"),
		CACHE_STORAGE:     os.Getenv("CACHE_STORAGE"),
		JWT_SECRET_KEY:    os.Getenv("JWT_SECRET_KEY"),
		APP_LOGO_URL:      "https://play-lh.googleusercontent.com/DTzWtkxfnKwFO3ruybY1SKjJQnLYeuK3KmQmwV5OQ3dULr5iXxeEtzBLceultrKTIUTr",
	}
}
