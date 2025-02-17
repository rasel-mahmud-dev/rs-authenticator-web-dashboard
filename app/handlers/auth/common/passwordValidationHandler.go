package common

import (
	"rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
	"rs/auth/app/services/hash"
	"rs/auth/app/utils"
	"time"
)

type PasswordValidationHandler struct {
	handlers.BaseHandler
}

func (h *PasswordValidationHandler) Handle(c *context.BaseContext) bool {
	loginRequest := c.LoginContext.LoginRequest
	user := c.User

	if user.Password == "" {
		response.Respond(c.ResponseWriter, statusCode.PASSWORD_NOT_CONFIGURED, "Password has not been configured for this account.", nil)
		return false
	}

	isMatchPassword := hash.Hash.VerifyHash(loginRequest.Password, user.Password)
	if !isMatchPassword {

		c.LoginContext.UserAuthAttempt = models.UserAuthAttempt{
			UserID:        user.ID,
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
		handler := InsertAuthFailedAttemptHandler{}
		handler.Handle(c)

		response.Respond(c.ResponseWriter, statusCode.INVALID_CREDENTIALS, "Invalid email or password", nil)
		return false
	}

	return h.HandleNext(c)
}
