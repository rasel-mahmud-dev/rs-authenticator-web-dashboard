package loginWithAuthenticator

import (
	"net/http"
	"rs/auth/app/dto"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
	"rs/auth/app/validators"
)

type RequestValidationHandler struct {
	handlers.BaseHandler
}

func (h *RequestValidationHandler) Handle(w http.ResponseWriter, r **http.Request) bool {
	payload := (*r).Context().Value("payload").(dto.AuthenticatorLoginRequestBody)
	err := validators.ValidateStruct(&dto.AuthenticatorLoginRequestBody{
		OtpCode: payload.OtpCode,
	})
	if err != nil {
		response.Respond(w, statusCode.REQUEST_VALIDATION_FAILED, err.Error(), nil)
		return false
	}
	return h.HandleNext(w, r)
}
