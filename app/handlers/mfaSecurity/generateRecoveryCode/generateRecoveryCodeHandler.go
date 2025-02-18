package generateRecoveryCode

import (
	"log"
	context2 "rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
	"rs/auth/app/repositories/recoveryCode"
	"rs/auth/app/utils"
	"time"
)

type GenerateRecoveryCodeHandler struct {
	handlers.BaseHandler
}

func (h *GenerateRecoveryCodeHandler) Handle(c *context2.BaseContext) bool {
	authSession := c.AuthSession

	isGeneratedNewRecoveryCode := c.TwoFaSecurityContext.IsGeneratedNewRecoveryCode
	recoveryCodes := c.TwoFaSecurityContext.RecoveryCodes
	if !isGeneratedNewRecoveryCode && len(recoveryCodes) > 0 {
		return h.HandleNext(c)
	}

	var codes []models.RecoveryCode

	for i := 0; i < 10; i++ {
		code := utils.RandomString(8)
		codes = append(codes, models.RecoveryCode{
			UserID:    authSession.UserId,
			Code:      &code,
			IsUsed:    false,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30 * 10),
		})
	}

	err := recoveryCode.RecoveryCodeRepository.InsertMultipleRecoveryCodes(codes)
	if err != nil {
		log.Fatalf("Error inserting batch: %v", err)
	}

	lastCodes, err := recoveryCode.RecoveryCodeRepository.GetLast10RecoveryCodes(authSession.UserId)
	if err != nil {
		utils.LoggerInstance.Error(err.Error())
	}

	if len(lastCodes) > 0 {
		c.TwoFaSecurityContext.RecoveryCodes = lastCodes
	}

	return h.HandleNext(c)
}
