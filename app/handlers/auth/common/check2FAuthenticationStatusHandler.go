package common

import (
	"rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/repositories"
	"rs/auth/app/response"
)

type Check2FAuthenticationStatusHandler struct {
	handlers.BaseHandler
}

func (h *Check2FAuthenticationStatusHandler) Handle(c *context.BaseContext) bool {
	user := c.User
	isEnabled := repositories.MfaSecurityTokenRepo.Is2FaEnabled(user.ID)
	if isEnabled {
		response.Respond(c.ResponseWriter, statusCode.OK, "Success", map[string]interface{}{
			"userId":     user.ID,
			"enabled2Fa": true,
		})
		return false
	}
	response.Respond(c.ResponseWriter, statusCode.INVALID_CREDENTIALS, "Invalid email or password", nil)
	return h.HandleNext(c)
}
