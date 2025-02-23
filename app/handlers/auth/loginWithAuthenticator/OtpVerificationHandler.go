package loginWithAuthenticator

import (
	context2 "rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/handlers/auth/common"
	"rs/auth/app/models"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/repositories"
	"rs/auth/app/response"
	"rs/auth/app/utils"
	"time"
)

type OtpVerificationHandler struct {
	handlers.BaseHandler
}

func (h *OtpVerificationHandler) Handle(c *context2.BaseContext) bool {
	payload := c.AuthenticatorLoginContext.RequestBody
	utils.LoggerInstance.Info("OtpVerificationHandler: ", payload)
	userId, err := repositories.MfaSecurityTokenRepo.VerifyMfaPasscode(payload.UserId, payload.OtpCode)
	if err != nil {
		c.LoginContext.UserAuthAttempt = models.UserAuthAttempt{
			UserID:        "",
			AttemptType:   "password_login",
			MFASecurityID: "",
			SecurityToken: "",
			IPAddress:     utils.GetUserIP(c.Request),
			UserAgent:     utils.GetUserAgent(c.Request),
			LastAttemptAt: time.Time{},
			IsSuccessful:  false,
			CreatedAt:     time.Time{},
			UpdatedAt:     time.Time{},
		}
		handler := common.InsertAuthFailedAttemptHandler{}
		handler.Handle(c)
		response.Respond(c.ResponseWriter, statusCode.INVALID_OTP, "Invalid otp code.", nil)
		return false
	}
	user, err := repositories.UserRepositoryInstance.GetUserById(userId)
	if err != nil || user == nil {
		utils.LoggerInstance.Info("User does not exist in database.")
		response.Respond(c.ResponseWriter, statusCode.INVALID_CREDENTIALS, "Invalid email or password", nil)
		return false
	}
	c.User = user
	return h.HandleNext(c)
}
