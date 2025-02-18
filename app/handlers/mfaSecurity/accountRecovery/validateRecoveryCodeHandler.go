package accountRecovery

import (
	"fmt"
	context2 "rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/repositories/recoveryCode"
	"rs/auth/app/response"
	"rs/auth/app/utils"
)

type ValidateRecoveryCodeHandler struct {
	handlers.BaseHandler
}

func (h *ValidateRecoveryCodeHandler) Handle(c *context2.BaseContext) bool {
	accountRecoveryCode := c.TwoFaSecurityContext.AccountRecoveryCode

	fmt.Println(accountRecoveryCode, "accountRecoveryCode")

	codeDetail, err := recoveryCode.RecoveryCodeRepository.GetValidRecoveryCode(accountRecoveryCode)
	if err != nil {
		utils.LoggerInstance.Error(err.Error())
		response.Respond(c.ResponseWriter, statusCode.INVALID_RECOVERY_CODE, "Internal error", nil)
		return false
	}

	if codeDetail.UserID == "" {
		response.Respond(c.ResponseWriter, statusCode.INVALID_RECOVERY_CODE, "Invalid recovery code", nil)
		return false
	}

	c.AccountRecoveryCodeRow = codeDetail

	return h.HandleNext(c)
}
