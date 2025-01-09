package login

import (
	"net/http"
	"rs/auth/app/dto"
	"rs/auth/app/response"
	"rs/auth/app/validators"
)

type RequestValidationHandler struct {
	BaseHandler
}

func (h *RequestValidationHandler) Handle(w http.ResponseWriter, r *http.Request) bool {
	loginRequest := r.Context().Value("loginRequest").(dto.LoginRequest)
	err := validators.ValidateLoginRequest(&loginRequest)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, err.Error(), nil)
		return false
	}
	return h.HandleNext(w, r)
}
