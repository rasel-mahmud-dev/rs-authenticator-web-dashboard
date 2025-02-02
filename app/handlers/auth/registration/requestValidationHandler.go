package registration

import (
	"fmt"
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

func (h *RequestValidationHandler) Handle(c context.BaseContext) bool {
	payload := c.RegistrationContext.Payload
	fmt.Println("Check request validation.")
	err := validators.ValidateStruct(&dto.RegisterRequestBody{
		Email:    payload.Email,
		Username: payload.Username,
		Password: payload.Password,
	})
	if err != nil {
		response.Respond(c.ResponseWriter, statusCode.REQUEST_VALIDATION_FAILED, err.Error(), nil)
		return false
	}
	c.Email = payload.Email
	return h.HandleNext(c)
}
