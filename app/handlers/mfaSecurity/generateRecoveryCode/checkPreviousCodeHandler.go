package generateRecoveryCode

import (
	"fmt"
	context2 "rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/repositories/recoveryCode"
	"rs/auth/app/utils"
)

type CheckPreviousCodeHandler struct {
	handlers.BaseHandler
}

func (h *CheckPreviousCodeHandler) Handle(c *context2.BaseContext) bool {
	userId := c.AuthSession.UserId
	isGeneratedNewRecoveryCode := c.TwoFaSecurityContext.IsGeneratedNewRecoveryCode

	if !isGeneratedNewRecoveryCode {
		fmt.Println("userId", userId)
		lastCodes, err := recoveryCode.RecoveryCodeRepository.GetLast10RecoveryCodes(userId)
		if err != nil {
			utils.LoggerInstance.Error(err.Error())
		}

		if len(lastCodes) == 0 {
			c.TwoFaSecurityContext.IsGeneratedNewRecoveryCode = false
		} else {
			c.TwoFaSecurityContext.RecoveryCodes = lastCodes
		}
	}

	return h.HandleNext(c)
}
