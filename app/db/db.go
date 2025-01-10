package db

import (
	"database/sql"
	"fmt"
	"log"
	"rs/auth/app/configs"
	"sync"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

var (
	dbInstance *sql.DB
	once       sync.Once
)

func GetDB() *sql.DB {
	once.Do(func() {
		config := configs.ConfigInstance()
		psqlInfo := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
			config.DATABASE_HOST,
			config.DATABASE_PORT,
			config.DATABASE_USER,
			config.DATABASE_PASSWORD,
			config.DATABASE_NAME,
		)
		var err error
		dbInstance, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Fatalf("Failed to open database connection: %v", err)
		}

		if err = dbInstance.Ping(); err != nil {
			log.Fatalf("Failed to ping database: %v", err)
		}
		dbInstance.SetMaxOpenConns(25)   // Maximum open connections
		dbInstance.SetMaxIdleConns(5)    // Maximum idle connections
		dbInstance.SetConnMaxLifetime(0) // No limit for connection lifetime
		log.Println("Database connection pool initialized")
	})

	return dbInstance
}

func CloseDB() {
	if dbInstance != nil {
		if err := dbInstance.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		} else {
			log.Println("Database connection pool closed")
		}
	}
}
