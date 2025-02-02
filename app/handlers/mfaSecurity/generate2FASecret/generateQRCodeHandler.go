package generate2FASecret

import (
	"encoding/base64"
	"fmt"
	"github.com/skip2/go-qrcode"
	context2 "rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
	"rs/auth/app/utils"
)

type GenerateQRCodeHandler struct {
	handlers.BaseHandler
}

func (h *GenerateQRCodeHandler) Handle(c *context2.BaseContext) bool {
	secretUrl := c.TwoFaSecurityContext.SecretUrl

	qrCodeData, err := qrcode.Encode(secretUrl, qrcode.Medium, 256)
	if err != nil {
		utils.LoggerInstance.Error("Failed to generate QR code")
		response.Respond(c.ResponseWriter, statusCode.INTERNAL_SERVER_ERROR, "QR code generate failed", nil)
		return false
	}

	qrBase64 := fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(qrCodeData))
	c.TwoFaSecurityContext.QrBase64 = qrBase64

	return h.HandleNext(c)
}
