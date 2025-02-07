package loginWithAuthenticator

import (
	context2 "rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/repositories"
	"rs/auth/app/response"
	"rs/auth/app/utils"
)

type OtpVerificationHandler struct {
	handlers.BaseHandler
}

func (h *OtpVerificationHandler) Handle(c *context2.BaseContext) bool {
	payload := c.AuthenticatorLoginContext.RequestBody
	utils.LoggerInstance.Info("OtpVerificationHandler: ", payload)
	userId, err := repositories.MfaSecurityTokenRepo.VerifyMfaPasscode(payload.OtpCode)
	if err != nil {
		response.Respond(c.ResponseWriter, statusCode.INVALID_OTP, "Invalid otp code.", nil)
		return false
	}
	userRepo := repositories.NewUserRepository()
	user, err := userRepo.GetUserById(userId)
	if err != nil || user == nil {
		utils.LoggerInstance.Info("User does not exist in database.")
		response.Respond(c.ResponseWriter, statusCode.INVALID_CREDENTIALS, "Invalid email or password", nil)
		return false
	}
	c.User = user
	return h.HandleNext(c)
}
