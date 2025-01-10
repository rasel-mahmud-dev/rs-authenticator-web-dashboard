package login

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

func (h *RequestValidationHandler) Handle(w http.ResponseWriter, r *http.Request) bool {
	loginRequest := r.Context().Value("loginRequest").(dto.LoginRequest)
	err := validators.ValidateStruct(&dto.LoginRequest{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	})
	if err != nil {
		response.Respond(w, statusCode.REQUEST_VALIDATION_FAILED, err.Error(), nil)
		return false
	}
	return h.HandleNext(w, r)
}
