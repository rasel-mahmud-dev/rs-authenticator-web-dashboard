package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"go/build"
	"log"
	"os"
	"path"
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
}

var config *Config
var once sync.Once

func init() {

	pkg, err := build.Import(".", ".", build.FindOnly)
	if err != nil {
		fmt.Println("Error:", err)
	}

	//p := path.Join(pkg.Dir, ".env")
	//fmt.Println(p)
	//data, err := os.ReadFile(p)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//os.Stdout.Write(data)

	err = godotenv.Load(path.Join(pkg.Dir, ".env"))
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Error loading .env file")
	}

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
		}
	})
}

func ConfigInstance() *Config {
	return config
}
