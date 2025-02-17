package generate2FASecret

import (
	"fmt"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"rs/auth/app/configs"
	"rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
	"rs/auth/app/utils"
)

type GenerateTotpSecretHandler struct {
	handlers.BaseHandler
}

func (h *GenerateTotpSecretHandler) Handle(c *context.BaseContext) bool {
	authSession := c.AuthSession

	codeName := fmt.Sprintf("RsAuth (%s)", authSession.Email)
	secret, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "RsAuthenticatorWeb2",
		AccountName: fmt.Sprintf("%s|%s", authSession.Email, configs.Config.APP_LOGO_URL),
		Period:      30,
		Digits:      otp.DigitsSix,
		Algorithm:   otp.AlgorithmSHA256,
	})
	if err != nil {
		utils.LoggerInstance.Error("Failed to generate secret")
		response.Respond(c.ResponseWriter, statusCode.INTERNAL_SERVER_ERROR, "QFailed to generate secret", nil)
		return false
	}

	c.TwoFaSecurityContext.CodeName = codeName
	c.TwoFaSecurityContext.SecretKey = secret.Secret()
	c.TwoFaSecurityContext.SecretUrl = secret.URL()

	return h.HandleNext(c)
}
