package login

import (
	"net/http"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
)

type ResponseHandler struct {
	handlers.BaseHandler
}

func (h *ResponseHandler) Handle(w http.ResponseWriter, r *http.Request) bool {
	user := r.Context().Value("user").(*models.User)
	authSession := r.Context().Value("authSession").(*models.AuthSession)
	token := r.Context().Value("token").(string)
	user.Password = ""
	response.Respond(w, statusCode.OK, "Login successful", struct {
		*models.User
		Token            string `json:"token"`
		SessionId        string `json:"sessionId"`
		IsRevokedSession bool   `json:"isRevokedSession"`
	}{
		User:             user,
		Token:            token,
		SessionId:        authSession.ID,
		IsRevokedSession: authSession.IsRevoked,
	})
	return false
}
