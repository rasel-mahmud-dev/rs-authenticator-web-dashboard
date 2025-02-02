package testUtils

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func (T *TestUtils) SendPostRequest(router *mux.Router, url string, requestBody map[string]interface{}, t *testing.T, tkn *string) *httptest.ResponseRecorder {
	body, _ := json.Marshal(requestBody)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	if tkn != nil {
		req.Header.Set("Authorization", "Bearer "+*tkn)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}
