package integration

import (
	"net/http"
	"net/http/httptest"
	"rs/auth/app/routes"
	"rs/auth/app/utils"
	testUtils "rs/auth/test/utils"
	"strconv"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func TestGenerateTwoFaSecretIntegration(t *testing.T) {
	SetupEnv()
	defer CleanupEnv()

	router := mux.NewRouter()
	routes.Init(router)

	t.Run("Should login successfully with valid credentials", func(t *testing.T) {
		email := strconv.Itoa(int(time.Now().UnixMilli())) + "test-@gmail.com"
		testUtilsInstance.Registration(router, email, t)

		payload := map[string]interface{}{
			"email":    email,
			"password": "123456",
		}

		rr := testUtilsInstance.SendPostRequest(router, "/api/v1/login", payload, t, nil)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expected status 200, got %v", status)
		}

		var response = testUtils.DecodeJSONResponse[map[string]interface{}](rr, t)
		token := utils.MapKey(response, "data", "token").(string)
		if token == "" {
			t.Errorf("Expected token, got empty string")
		}

		req, err := http.NewRequest(http.MethodGet, "/api/v1/generate-2fa-secret", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Authorization", "Bearer "+token)
		rr = httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		response = testUtils.DecodeJSONResponse[map[string]interface{}](rr, t)
		id := utils.MapKey(response, "data", "id").(string)
		//secret := utils.MapKey(response, "data", "secret").(string)
		if id == "" {
			t.Errorf("Expected id, got empty string")
		}

		payload = map[string]interface{}{
			"id":          id,
			"provider":    "Google",
			"isCompleted": true,
		}
		rr = testUtilsInstance.SendPostRequest(router, "/api/v1/generate-2fa-secret", payload, t, &token)
		response = testUtils.DecodeJSONResponse[map[string]interface{}](rr, t)
		statusCode := utils.MapKey(response, "statusCode").(string)

		if statusCode != "SUCCESS-000" {
			t.Errorf("Expected status code SUCCESS-000, got %v", statusCode)
		}

	})

}
