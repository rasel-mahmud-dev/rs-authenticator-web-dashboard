package accountRecovery

import (
	context2 "rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
	"rs/auth/app/utils"
)

type ResponseHandler struct {
	handlers.BaseHandler
}

func (h *ResponseHandler) Handle(c *context2.BaseContext) bool {
	recoveryCodes := c.TwoFaSecurityContext.RecoveryCodes
	utils.LoggerInstance.Error("Failed to generated or fetch account recovery codes.")
	response.Respond(c.ResponseWriter, statusCode.OK, "Success", recoveryCodes)
	return false
}
