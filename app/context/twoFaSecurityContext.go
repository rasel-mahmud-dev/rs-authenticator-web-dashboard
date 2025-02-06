package context

import "rs/auth/app/dto"

type TwoFaSecurityContext struct {
	SecretKey       string
	SecretUrl       string
	CodeName        string
	QrBase64        string
	GenerateMfaBody dto.GenerateMfaQRRequestPayload
}
