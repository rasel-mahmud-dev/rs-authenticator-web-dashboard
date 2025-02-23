package accountRecovery

import (
	"rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/repositories"
	"rs/auth/app/repositories/recoveryCode"
	"rs/auth/app/response"
)

type PreparedContextForCreateNewSession struct {
	handlers.BaseHandler
}

func (h *PreparedContextForCreateNewSession) Handle(c *context.BaseContext) bool {
	recoveryCodeRow := c.AccountRecoveryCodeRow
	if recoveryCodeRow.UserID == "" {
		response.Respond(c.ResponseWriter, statusCode.INVALID_RECOVERY_CODE, "Invalid recovery code", nil)
		return false
	}

	user, err := repositories.UserRepositoryInstance.GetUserById(recoveryCodeRow.UserID)
	if err != nil {
		response.Respond(c.ResponseWriter, statusCode.INVALID_RECOVERY_CODE, "Internal error user fetch failed from database.", nil)
		return false
	}

	err = recoveryCode.RecoveryCodeRepository.MakeInvalidRecoveryCodeById(recoveryCodeRow.ID)
	if err != nil {
		response.Respond(c.ResponseWriter, statusCode.INVALID_RECOVERY_CODE, "Account Recovery code checking failed.", nil)
		return false
	}

	c.AuthMethod = "recovery_code"
	c.User = user
	return h.HandleNext(c)
}
