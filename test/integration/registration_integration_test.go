package integration

import (
	"github.com/gorilla/mux"
	"net/http"
	"rs/auth/app/routes"
	"rs/auth/app/utils"
	"rs/auth/test/utils"
	"strconv"
	"testing"
	"time"
)

var testUtilsInstance = &testUtils.TestUtils{}

func TestRegistrationIntegration(t *testing.T) {
	SetupEnv()
	defer CleanupEnv()

	router := mux.NewRouter()
	routes.Init(router)

	t.Run("Should successfully create an account.", func(t *testing.T) {
		email := strconv.Itoa(int(time.Now().UnixMilli())) + "test-@gmail.com"
		payload := map[string]interface{}{
			"username": "admin",
			"email":    email,
			"password": "123456",
		}
		rr := testUtilsInstance.SendPostRequest(router, "/api/v1/registration", payload, t)
		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("Expected status 201, got %v", status)
		}

		var response = testUtils.DecodeJSONResponse[map[string]interface{}](rr, t)
		email = utils.MapKey(response, "data", "email").(string)
		if email != payload["email"] {
			t.Errorf("Expected status '%s', got '%s'", "expectedMessage", response["status"])
		}

	})
}
