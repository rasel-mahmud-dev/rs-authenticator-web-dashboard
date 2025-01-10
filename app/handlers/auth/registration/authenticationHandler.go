package registration

import (
	"net/http"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
	"rs/auth/app/response"
)

type AuthenticationHandler struct {
	handlers.BaseHandler
}

func (h *AuthenticationHandler) Handle(w http.ResponseWriter, r *http.Request) bool {
	user := r.Context().Value("user").(*models.User)
	user.Password = ""
	response.Respond(w, http.StatusOK, "Account successful created", user)
	return false
}
