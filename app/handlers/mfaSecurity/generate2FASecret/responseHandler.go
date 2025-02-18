package generate2FASecret

import (
	context2 "rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
	"time"
)

type ResponseHandler struct {
	handlers.BaseHandler
}

func (h *ResponseHandler) Handle(c *context2.BaseContext) bool {
	authSession := c.AuthSession
	qrBase64 := c.TwoFaSecurityContext.QrBase64
	codeName := c.TwoFaSecurityContext.CodeName
	secretKey := c.TwoFaSecurityContext.SecretKey

	mfaToken := models.MfaSecurityToken{
		UserID:    authSession.UserId,
		CodeName:  &codeName,
		Secret:    secretKey,
		QrCodeURL: &qrBase64,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		LinkedAt:  nil,
		AppName:   "RsAuth",
	}
	
	response.Respond(c.ResponseWriter, statusCode.OK, "Success", mfaToken)
	return false
}
