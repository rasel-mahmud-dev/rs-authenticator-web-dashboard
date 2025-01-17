package loginWithAuthenticator

import (
	"context"
	"fmt"
	"net/http"
	"rs/auth/app/dto"
	"rs/auth/app/handlers"
	"rs/auth/app/repositories"
)

type OtpVerificationHandler struct {
	handlers.BaseHandler
}

func (h *OtpVerificationHandler) Handle(w http.ResponseWriter, r **http.Request) bool {
	payload := (*r).Context().Value("payload").(dto.AuthenticatorLoginRequestBody)

	userId, err := repositories.MfaSecurityTokenRepo.VerifyMfaPasscode(payload.OtpCode)
	if err != nil {
		return false
	}
	fmt.Println(userId)
	ctx := context.WithValue((*r).Context(), "payload", payload)
	*r = (*r).WithContext(ctx)
	return h.HandleNext(w, r)
}
