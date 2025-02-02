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

func TestAuthVerifyIntegration(t *testing.T) {
	SetupEnv()
	defer CleanupEnv()

	router := mux.NewRouter()
	routes.Init(router)

	email := strconv.Itoa(int(time.Now().UnixMilli())) + "test-@gmail.com"
	testUtilsInstance.Registration(router, email, t)

	t.Run("Should successfully verify user with valid token", func(t *testing.T) {
		loginRequest := map[string]interface{}{
			"email":    email,
			"password": "123456",
		}

		rr := testUtilsInstance.SendPostRequest(router, "/api/v1/login", loginRequest, t)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expected status 200, got %v", status)
		}

		var response = testUtils.DecodeJSONResponse[map[string]interface{}](rr, t)
		token := utils.MapKey(response, "data", "token").(string)
		if token == "" {
			t.Errorf("Expected sessionId, got empty string")
		}

		req, err := http.NewRequest(http.MethodGet, "/api/v1/verify", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Authorization", "Bearer "+token)
		rr = httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		response = testUtils.DecodeJSONResponse[map[string]interface{}](rr, t)
		sessionId := utils.MapKey(response, "data", "sessionId").(string)
		if sessionId == "" {
			t.Errorf("Expected sessionId, got empty string")
		}
	})
}
