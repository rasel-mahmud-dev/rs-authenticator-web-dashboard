package generate2FASecret

import (
	context2 "rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/repositories"
	"rs/auth/app/response"
	"rs/auth/app/utils"
)

type ResponseHandler struct {
	handlers.BaseHandler
}

func (h *ResponseHandler) Handle(c *context2.BaseContext) bool {
	authSession := c.AuthSession
	qrBase64 := c.TwoFaSecurityContext.QrBase64
	codeName := c.TwoFaSecurityContext.CodeName
	secretKey := c.TwoFaSecurityContext.SecretKey

	backupCodes := utils.GenerateBackupCodes(12)

	mfaToken, err := repositories.MfaSecurityTokenRepo.InsertMfaSecurityToken(models.MfaSecurityToken{
		UserID:        authSession.UserId,
		Secret:        secretKey,
		CodeName:      &codeName,
		RecoveryCodes: backupCodes,
		QrCodeURL:     &qrBase64,
		DeviceInfo:    nil,
	})

	if err != nil {
		utils.LoggerInstance.Error("Failed to save 2FA secret")
		response.Respond(c.ResponseWriter, statusCode.INTERNAL_SERVER_ERROR, "Failed to save 2FA secret", nil)
		return false
	}

	response.Respond(c.ResponseWriter, statusCode.OK, "Success", mfaToken)
	return false
}
