package testUtils

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func (T *TestUtils) Registration(router *mux.Router, email string, t *testing.T) map[string]interface{} {
	body, _ := json.Marshal(map[string]interface{}{
		"username": "admin",
		"email":    email,
		"password": "123456",
	})
	req, err := http.NewRequest(http.MethodPost, "/api/v1/registration", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Expected status 201, got %v", status)
	}

	return DecodeJSONResponse[map[string]interface{}](rr, t)
}
