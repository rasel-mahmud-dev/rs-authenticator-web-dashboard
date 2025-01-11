package generate2FASecret

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/skip2/go-qrcode"
	"net/http"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
	"rs/auth/app/utils"
)

type GenerateQRCodeHandler struct {
	handlers.BaseHandler
}

func (h *GenerateQRCodeHandler) Handle(w http.ResponseWriter, r **http.Request) bool {
	secretUrl := (*r).Context().Value("secretUrl").(string)

	qrCodeData, err := qrcode.Encode(secretUrl, qrcode.Medium, 256)
	if err != nil {
		utils.LoggerInstance.Error("Failed to generate QR code")
		response.Respond(w, statusCode.INTERNAL_SERVER_ERROR, "QR code generate failed", nil)
		return false
	}

	qrBase64 := fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(qrCodeData))
	ctx := context.WithValue((*r).Context(), "qrBase64", qrBase64)
	*r = (*r).WithContext(ctx)

	return h.HandleNext(w, r)
}
