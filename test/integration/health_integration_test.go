package integration

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"rs/auth/app/routes"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	SetupEnv()
	defer CleanupEnv()

	router := mux.NewRouter()
	routes.Init(router)

	req, err := http.NewRequest(http.MethodGet, "/api/health", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status 200, got %v", status)
	}

	var response map[string]string
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err.Error())
	}

	expected := "Healthy"
	if response["status"] != expected {
		t.Errorf("Expected status '%s', got '%s'", expected, response["status"])
	}
}
