package generateRecoveryCode

import (
	"rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
	"strings"
)

type PreparedContextState struct {
	handlers.BaseHandler
}

func (h *PreparedContextState) Handle(c *context.BaseContext) bool {

	isNew := strings.ToLower(c.Request.URL.Query().Get("isNew")) == "true"
	c.TwoFaSecurityContext.IsGeneratedNewRecoveryCode = isNew

	authSession := c.Request.Context().Value("authSession").(*models.AuthSession)
	c.AuthSession = authSession
	return h.HandleNext(c)
}
