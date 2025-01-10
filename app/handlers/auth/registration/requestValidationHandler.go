package registration

import (
	"net/http"
	"rs/auth/app/dto"
	"rs/auth/app/handlers"
	"rs/auth/app/response"
	"rs/auth/app/validators"
)

type RequestValidationHandler struct {
	handlers.BaseHandler
}

func (h *RequestValidationHandler) Handle(w http.ResponseWriter, r *http.Request) bool {
	payload := r.Context().Value("payload").(dto.RegisterRequestBody)
	err := validators.ValidateStruct(&dto.RegisterRequestBody{
		Email:    payload.Email,
		Username: payload.Username,
		Password: payload.Password,
	})
	if err != nil {
		response.Respond(w, http.StatusBadRequest, err.Error(), nil)
		return false
	}
	return h.HandleNext(w, r)
}
