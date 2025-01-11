package generate2FASecret

import (
	"net/http"
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

func (h *ResponseHandler) Handle(w http.ResponseWriter, r **http.Request) bool {
	authSession := (*r).Context().Value("authSession").(*models.AuthSession)
	qrBase64 := (*r).Context().Value("qrBase64").(string)
	codeName := (*r).Context().Value("codeName").(string)
	secretKey := (*r).Context().Value("secretKey").(string)

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
		response.Respond(w, statusCode.INTERNAL_SERVER_ERROR, "Failed to save 2FA secret", nil)
		return false
	}

	response.Respond(w, statusCode.OK, "Success", mfaToken)
	return false
}
