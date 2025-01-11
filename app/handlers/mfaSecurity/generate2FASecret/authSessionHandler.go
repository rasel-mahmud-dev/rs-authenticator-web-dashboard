package generate2FASecret

import (
	"net/http"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
)

type AuthSessionHandler struct {
	handlers.BaseHandler
}

func (h *AuthSessionHandler) Handle(w http.ResponseWriter, r **http.Request) bool {
	authSession := (*r).Context().Value("authSession").(*models.AuthSession)

	if authSession == nil {
		response.Respond(w, statusCode.UNAUTHORIZED, "UNAUTHORIZED", nil)
		return false
	}

	return h.HandleNext(w, r)
}
