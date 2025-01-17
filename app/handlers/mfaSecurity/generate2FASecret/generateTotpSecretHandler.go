package generate2FASecret

import (
	"context"
	"fmt"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"net/http"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
	"rs/auth/app/utils"
)

type GenerateTotpSecretHandler struct {
	handlers.BaseHandler
}

func (h *GenerateTotpSecretHandler) Handle(w http.ResponseWriter, r **http.Request) bool {
	authSession := (*r).Context().Value("authSession").(*models.AuthSession)

	codeName := fmt.Sprintf("RsAuth (%s)", authSession.Email)
	secret, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "RsAuth",
		AccountName: authSession.Email,
		Period:      30,
		Digits:      otp.DigitsSix,
		Algorithm:   otp.AlgorithmSHA1,
	})
	if err != nil {
		utils.LoggerInstance.Error("Failed to generate secret")
		response.Respond(w, statusCode.INTERNAL_SERVER_ERROR, "QFailed to generate secret", nil)
		return false
	}

	ctx := context.WithValue((*r).Context(), "secretKey", secret.Secret())
	ctx = context.WithValue(ctx, "secretUrl", secret.URL())
	ctx = context.WithValue(ctx, "codeName", codeName)
	*r = (*r).WithContext(ctx)

	return h.HandleNext(w, r)
}
