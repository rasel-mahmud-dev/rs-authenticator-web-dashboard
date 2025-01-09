package integration

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func SetupEnv() {
	err := godotenv.Load("../../.env.test")
	if err != nil {
		log.Fatalf("Error loading .env file for tests: %v", err)
	}
}

func CleanupEnv() {
	os.Clearenv()
}

func TestMain(m *testing.M) {
	SetupEnv()
	code := m.Run()
	CleanupEnv()
	os.Exit(code)
}
