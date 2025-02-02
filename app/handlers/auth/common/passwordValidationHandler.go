package authCommon

import (
	"net/http"
	"rs/auth/app/dto"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
	"rs/auth/app/services/hash"
)

type PasswordValidationHandler struct {
	handlers.BaseHandler
}

func (h *PasswordValidationHandler) Handle(w http.ResponseWriter, r **http.Request) bool {
	loginRequest := (*r).Context().Value("loginRequest").(dto.LoginRequest)
	user := (*r).Context().Value("user").(*models.User)

	if user.Password == "" {
		response.Respond(w, statusCode.PASSWORD_NOT_CONFIGURED, "Password has not been configured for this account.", nil)
		return false
	}

	isMatchPassword := hash.Hash.VerifyHash(loginRequest.Password, user.Password)
	if !isMatchPassword {
		response.Respond(w, statusCode.INVALID_CREDENTIALS, "Invalid email or password", nil)
		return false
	}

	return h.HandleNext(w, r)
}
