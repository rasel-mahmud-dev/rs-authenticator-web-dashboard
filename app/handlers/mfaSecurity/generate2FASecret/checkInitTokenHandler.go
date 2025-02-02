package generate2FASecret

import (
	context2 "rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/repositories"
	"rs/auth/app/response"
	"rs/auth/app/utils"
)

type CheckInitTokenHandler struct {
	handlers.BaseHandler
}

func (h *CheckInitTokenHandler) Handle(c *context2.BaseContext) bool {
	authSession := c.AuthSession

	initToken, err := repositories.MfaSecurityTokenRepo.GetLastInit(authSession.UserId)
	if err != nil {
		utils.LoggerInstance.Error(err.Error())
	}

	if initToken != nil {
		utils.LoggerInstance.Error("Already  exists previous token.")
		response.Respond(c.ResponseWriter, statusCode.OK, "Success", initToken)
		return false
	}

	return h.HandleNext(c)
}
