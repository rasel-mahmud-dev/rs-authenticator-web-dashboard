package twoFactor

import (
	"encoding/base64"
	"fmt"
	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
	"net/http"
	"rs/auth/app/models"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/repositories"
	"rs/auth/app/response"
	"rs/auth/app/utils"
)

func Generate2FASecret(w http.ResponseWriter, r *http.Request) {

	authSession := (*r).Context().Value("authSession").(*models.AuthSession)

	if authSession == nil {
		response.Respond(w, statusCode.UNAUTHORIZED, "UNAUTHORIZED", nil)
		return
	}

	secret, err := totp.Generate(totp.GenerateOpts{
		Issuer:      fmt.Sprintf("RsAuth (%s)", authSession.Email),
		AccountName: authSession.Email,
	})

	if err != nil {
		utils.LoggerInstance.Error("Failed to generate secret")
		response.Respond(w, statusCode.INTERNAL_SERVER_ERROR, "QFailed to generate secret", nil)
		return
	}

	secretKey := secret.Secret()

	qrCodeData, err := qrcode.Encode(secret.URL(), qrcode.Medium, 256)
	if err != nil {
		utils.LoggerInstance.Error("Failed to generate QR code")
		response.Respond(w, statusCode.INTERNAL_SERVER_ERROR, "QR code generate failed", nil)
		return
	}

	qrBase64 := fmt.Sprintf("data:image/png;base64,%s", toBase64(qrCodeData))
	backupCodes := utils.GenerateBackupCodes(12)

	mfaToken, err := repositories.MfaSecurityTokenRepo.InsertMfaSecurityToken(models.MfaSecurityToken{
		UserID:        authSession.UserId,
		Secret:        secretKey,
		RecoveryCodes: backupCodes,
		QrCodeURL:     &qrBase64,
		DeviceInfo:    nil,
	})

	if err != nil {
		utils.LoggerInstance.Error("Failed to save 2FA secret")
		response.Respond(w, statusCode.INTERNAL_SERVER_ERROR, "Failed to save 2FA secret", nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response.Respond(w, statusCode.OK, "Success", mfaToken)
}

func toBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
