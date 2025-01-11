package twoFactor

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
	"net/http"
	"rs/auth/app/dto"
	"rs/auth/app/models"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/repositories"
	"rs/auth/app/response"
	"rs/auth/app/utils"
	"time"
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

func Finalize2FASecret(w http.ResponseWriter, r *http.Request) {

	authSession := (*r).Context().Value("authSession").(*models.AuthSession)

	if authSession == nil {
		response.Respond(w, statusCode.UNAUTHORIZED, "UNAUTHORIZED", nil)
		return
	}

	var body dto.Completed2FASecretBody
	err := json.NewDecoder((*r).Body).Decode(&body)
	if err != nil {
		response.Respond(w, statusCode.INVALID_JSON_FORMAT, "Invalid JSON format", nil)
		return
	}

	_, err = repositories.MfaSecurityTokenRepo.GetById(body.Id, authSession.UserId)
	if err != nil {
		response.Respond(w, statusCode.INVALID_JSON_FORMAT, "Setup session destroyed", nil)
		return
	}

	err = repositories.MfaSecurityTokenRepo.UpdateMfaSecurityToken(models.MfaSecurityToken{
		UserID:    authSession.UserId,
		AppName:   body.AppName,
		IsActive:  body.IsCompleted,
		ID:        body.Id,
		UpdatedAt: time.Now(),
	})

	if err != nil {
		utils.LoggerInstance.Error(err.Error())
		response.Respond(w, statusCode.INTERNAL_SERVER_ERROR, "Unable to completed authenticator app setup", nil)
		return
	}

	response.Respond(w, statusCode.OK, "OK", nil)
}

func toBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
