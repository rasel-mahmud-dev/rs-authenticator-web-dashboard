package registration

import (
	"net/http"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
)

type AuthenticationHandler struct {
	handlers.BaseHandler
}

func (h *AuthenticationHandler) Handle(w http.ResponseWriter, r *http.Request) bool {
	user := r.Context().Value("user").(*models.User)
	user.Password = ""
	response.Respond(w, statusCode.OK, "Account successful created", user)
	return false
}
