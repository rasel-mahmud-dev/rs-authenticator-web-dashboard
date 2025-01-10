package login

import (
	"context"
	"encoding/json"
	"net/http"
	"rs/auth/app/dto"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
)

type JSONValidationHandler struct {
	handlers.BaseHandler
}

func (h *JSONValidationHandler) Handle(w http.ResponseWriter, r *http.Request) bool {
	var loginRequest dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		response.Respond(w, statusCode.INVALID_JSON_FORMAT, "Invalid JSON format", nil)
		return false
	}

	ctx := context.WithValue(r.Context(), "loginRequest", loginRequest)
	r = r.WithContext(ctx)
	return h.HandleNext(w, r)
}
