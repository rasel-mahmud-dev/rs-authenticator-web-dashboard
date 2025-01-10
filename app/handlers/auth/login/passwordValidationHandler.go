package login

import (
	"net/http"
	"rs/auth/app/dto"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
	"rs/auth/app/response"
)

type PasswordValidationHandler struct {
	handlers.BaseHandler
}

func (h *PasswordValidationHandler) Handle(w http.ResponseWriter, r *http.Request) bool {
	loginRequest := r.Context().Value("loginRequest").(dto.LoginRequest)
	user := r.Context().Value("user").(*models.User)

	if loginRequest.Password != user.Password {
		response.Respond(w, http.StatusUnauthorized, "Invalid username or password", nil)
		return false
	}

	response.Respond(w, http.StatusOK, "Login successful", nil)
	return false
}
