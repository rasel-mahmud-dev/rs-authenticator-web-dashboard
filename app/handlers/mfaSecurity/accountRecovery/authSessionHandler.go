package accountRecovery

import (
	"rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
)

type AuthSessionHandler struct {
	handlers.BaseHandler
}

func (h *AuthSessionHandler) Handle(c *context.BaseContext) bool {
	authSession := c.AuthSession

	if authSession == nil {
		response.Respond(c.ResponseWriter, statusCode.UNAUTHORIZED, "UNAUTHORIZED", nil)
		return false
	}

	return h.HandleNext(c)
}
