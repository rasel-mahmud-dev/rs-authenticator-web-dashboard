package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rs/auth/internal/app/handlers/auth"
	"testing"

	"github.com/gorilla/mux"
)

func TestLoginIntegration(t *testing.T) {
	SetupEnv()
	defer CleanupEnv()

	router := mux.NewRouter()
	router.HandleFunc("/login", authHandler.LoginHandler).Methods("POST")

	t.Run("Valid credentials", func(t *testing.T) {
		loginRequest := map[string]string{
			"email":    "rasel.mahmud.dev@gmail.com",
			"password": "123456",
		}
		loginBody, _ := json.Marshal(loginRequest)
		req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(loginBody))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(authHandler.LoginHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expected status 200, got %v", status)
		}

		var response map[string]string
		if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}
		expected := "Login successful"
		if response["message"] != expected {
			t.Errorf("Expected message '%s', got '%s'", expected, response["message"])
		}
	})

	//t.Run("Invalid credentials", func(t *testing.T) {
	//	loginRequest := map[string]string{
	//		"username": "admin",
	//		"password": "wrongpassword", // Invalid credentials
	//	}
	//	loginBody, _ := json.Marshal(loginRequest)
	//	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(loginBody))
	//	if err != nil {
	//		t.Fatalf("Failed to create request: %v", err)
	//	}
	//
	//	rr := httptest.NewRecorder()
	//	handler := http.HandlerFunc(authHandler.LoginHandler)
	//	handler.ServeHTTP(rr, req)
	//
	//	if status := rr.Code; status != http.StatusUnauthorized {
	//		t.Errorf("Expected status 401, got %v", status)
	//	}
	//
	//	var response map[string]string
	//	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
	//		t.Fatalf("Failed to decode response: %v", err)
	//	}
	//	expected := "Invalid username or password"
	//	if response["message"] != expected {
	//		t.Errorf("Expected message '%s', got '%s'", expected, response["message"])
	//	}
	//})
}
