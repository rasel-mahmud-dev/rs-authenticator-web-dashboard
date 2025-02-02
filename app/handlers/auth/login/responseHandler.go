package login

import (
	context2 "rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
)

type ResponseHandler struct {
	handlers.BaseHandler
}

func (h *ResponseHandler) Handle(c *context2.BaseContext) bool {
	user := c.User
	authSession := c.AuthSession
	token := c.AccessToken
	user.Password = ""
	response.Respond(c.ResponseWriter, statusCode.OK, "Login successful", struct {
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
