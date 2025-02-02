package login

import (
	"encoding/json"
	context2 "rs/auth/app/context"
	"rs/auth/app/dto"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
)

type JSONValidationHandler struct {
	handlers.BaseHandler
}

func (h *JSONValidationHandler) Handle(c context2.BaseContext) bool {
	var loginRequest dto.LoginRequest
	err := json.NewDecoder((*c.Request).Body).Decode(&loginRequest)
	if err != nil {
		response.Respond(c.ResponseWriter, statusCode.INVALID_JSON_FORMAT, "Invalid JSON format", nil)
		return false
	}

	c.LoginContext.LoginRequest = loginRequest

	return h.HandleNext(c)
}
