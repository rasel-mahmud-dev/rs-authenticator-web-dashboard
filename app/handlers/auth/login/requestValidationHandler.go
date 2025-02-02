package login

import (
	"rs/auth/app/context"
	"rs/auth/app/dto"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
	"rs/auth/app/validators"
)

type RequestValidationHandler struct {
	handlers.BaseHandler
}

func (h *RequestValidationHandler) Handle(c *context.BaseContext) bool {
	loginRequest := c.LoginContext.LoginRequest
	err := validators.ValidateStruct(&dto.LoginRequest{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	})
	if err != nil {
		response.Respond(c.ResponseWriter, statusCode.REQUEST_VALIDATION_FAILED, err.Error(), nil)
		return false
	}
	return h.HandleNext(c)
}
