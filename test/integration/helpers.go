package integration

import (
	"os"
	"testing"
)

func SetupEnv() {
	os.Setenv("PORT", "8080")
}

func CleanupEnv() {
	os.Unsetenv("PORT")
}

func TestMain(m *testing.M) {
	SetupEnv()
	code := m.Run()
	CleanupEnv()
	os.Exit(code)
}
