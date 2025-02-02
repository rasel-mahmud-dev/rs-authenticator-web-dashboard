package loginWithAuthenticator

import (
	"context"
	"net/http"
	"rs/auth/app/dto"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/repositories"
	"rs/auth/app/response"
	"rs/auth/app/utils"
)

type OtpVerificationHandler struct {
	handlers.BaseHandler
}

func (h *OtpVerificationHandler) Handle(w http.ResponseWriter, r **http.Request) bool {
	payload := (*r).Context().Value("payload").(dto.AuthenticatorLoginRequestBody)

	userId, err := repositories.MfaSecurityTokenRepo.VerifyMfaPasscode(payload.OtpCode)
	if err != nil {
		response.Respond(w, statusCode.INVALID_OTP, "Invalid otp code.", nil)
		return false
	}
	userRepo := repositories.NewUserRepository()
	user, err := userRepo.GetUserById(userId)
	if err != nil || user == nil {
		utils.LoggerInstance.Info("User does not exist in database.")
		response.Respond(w, statusCode.INVALID_CREDENTIALS, "Invalid email or password", nil)
		return false
	}

	ctx := context.WithValue((*r).Context(), "user", user)
	*r = (*r).WithContext(ctx)

	return h.HandleNext(w, r)
}
