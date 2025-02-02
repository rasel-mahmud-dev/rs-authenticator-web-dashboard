package integration

import (
	"os"
	"testing"
)

func SetupEnv() {
	 
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
