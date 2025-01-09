package login

import (
	"context"
	"encoding/json"
	"net/http"
	"rs/auth/app/dto"
	"rs/auth/app/response"
)

type JSONValidationHandler struct {
	BaseHandler
}

func (h *JSONValidationHandler) Handle(w http.ResponseWriter, r *http.Request) bool {
	var loginRequest dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, "Invalid JSON format", nil)
		return false
	}
	ctx := context.WithValue(r.Context(), "loginRequest", loginRequest)
	r = r.WithContext(ctx)
	return h.HandleNext(w, r)
}
