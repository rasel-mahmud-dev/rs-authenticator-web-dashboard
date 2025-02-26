package common

import (
	"rs/auth/app/configs"
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

	apiSecret := c.Request.Header.Get("API_SECRET")
	if apiSecret == configs.Config.MOBILE_API_SECRET {
		return h.HandleNext(c)
	}

	isEnabled := repositories.MfaSecurityTokenRepo.Is2FaEnabled(user.ID)
	if isEnabled {
		response.Respond(c.ResponseWriter, statusCode.OK, "Success", map[string]interface{}{
			"userId":     user.ID,
			"enabled2Fa": true,
		})
		return false
	}
	return h.HandleNext(c)
}
