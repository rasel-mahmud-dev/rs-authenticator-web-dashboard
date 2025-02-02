package integration

import (
	"net/http"
	"rs/auth/app/routes"
	"rs/auth/app/utils"
	testUtils "rs/auth/test/utils"
	"strconv"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func TestLoginIntegration(t *testing.T) {
	SetupEnv()
	defer CleanupEnv()

	router := mux.NewRouter()
	routes.Init(router)

	// Register user before running login tests
	email := strconv.Itoa(int(time.Now().UnixMilli())) + "test-@gmail.com"
	testUtilsInstance.Registration(router, email, t)

	t.Run("Should login successfully with valid credentials", func(t *testing.T) {
		loginRequest := map[string]interface{}{
			"email":    email,
			"password": "123456",
		}

		rr := testUtilsInstance.SendPostRequest(router, "/api/v1/login", loginRequest, t)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expected status 200, got %v", status)
		}

		var response = testUtils.DecodeJSONResponse[map[string]interface{}](rr, t)
		sessionId := utils.MapKey(response, "data", "sessionId").(string)
		if sessionId == "" {
			t.Errorf("Expected sessionId, got empty string")
		}
	})

	t.Run("Should return user not found.", func(t *testing.T) {
		email := strconv.Itoa(int(time.Now().UnixMilli())) + "test-@gmail.com"
		loginRequest := map[string]interface{}{
			"email":    email,
			"password": "123456",
		}

		rr := testUtilsInstance.SendPostRequest(router, "/api/v1/login", loginRequest, t)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %v", status)
		}

		var response = testUtils.DecodeJSONResponse[map[string]interface{}](rr, t)
		statusCode := utils.MapKey(response, "statusCode").(string)
		if statusCode != "ERR-001" {
			t.Errorf("Expected ERR-001, for user not found, got %v", statusCode)
		}
	})

	t.Run("Should return incorrect password error", func(t *testing.T) {
		// Using the email registered in the first test
		loginRequest := map[string]interface{}{
			"email":    email,
			"password": "wrongpassword",
		}

		rr := testUtilsInstance.SendPostRequest(router, "/api/v1/login", loginRequest, t)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %v", status)
		}

		var response = testUtils.DecodeJSONResponse[map[string]interface{}](rr, t)
		statusCode := utils.MapKey(response, "statusCode").(string)
		if statusCode != "ERR-001" {
			t.Errorf("Expected ERR-001 for incorrect password, got %v", statusCode)
		}
	})

	t.Run("Should return missing password error", func(t *testing.T) {
		loginRequest := map[string]interface{}{
			"email": email,
		}

		rr := testUtilsInstance.SendPostRequest(router, "/api/v1/login", loginRequest, t)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %v", status)
		}

		var response = testUtils.DecodeJSONResponse[map[string]interface{}](rr, t)
		statusCode := utils.MapKey(response, "statusCode").(string)
		if statusCode != "ERR-023" {
			t.Errorf("Expected ERR-023 for missing password, got %v", statusCode)
		}
	})
}
