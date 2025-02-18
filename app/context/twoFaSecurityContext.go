package context

import (
	"rs/auth/app/dto"
	"rs/auth/app/models"
)

type TwoFaSecurityContext struct {
	SecretKey                  string
	SecretUrl                  string
	CodeName                   string
	QrBase64                   string
	GenerateMfaBody            dto.GenerateMfaQRRequestPayload
	IsGeneratedNewRecoveryCode bool
	RecoveryCodes              []models.RecoveryCode
	AccountRecoveryCode        string
}
