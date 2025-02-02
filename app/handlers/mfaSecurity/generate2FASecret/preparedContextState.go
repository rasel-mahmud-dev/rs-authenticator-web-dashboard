package generate2FASecret

import (
	"rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
)

type PreparedContextState struct {
	handlers.BaseHandler
}

func (h *PreparedContextState) Handle(c *context.BaseContext) bool {
	authSession := c.Request.Context().Value("authSession").(*models.AuthSession)
	c.AuthSession = authSession
	return h.HandleNext(c)
}
