package testUtils

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func DecodeJSONResponse[T any](rr *httptest.ResponseRecorder, t *testing.T) T {
	var response T
	err := json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err.Error())
	}
	return response
}
