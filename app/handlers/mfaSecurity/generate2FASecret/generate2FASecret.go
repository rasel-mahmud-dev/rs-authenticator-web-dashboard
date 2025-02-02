package generate2FASecret

import (
	"net/http"
	context2 "rs/auth/app/context"
)

func Generate2FASecretHandler(w http.ResponseWriter, r *http.Request) {

	c := context2.BaseContext{
		RegistrationContext:  context2.RegistrationContext{},
		ResponseWriter:       w,
		Request:              r,
		Email:                "",
		User:                 nil,
		AuthSession:          nil,
		TwoFaSecurityContext: context2.TwoFaSecurityContext{},
	}

	authSessionHandler := &AuthSessionHandler{}

	chain := authSessionHandler
	chain.SetNext(&CheckInitTokenHandler{}).
		SetNext(&GenerateTotpSecretHandler{}).
		SetNext(&GenerateQRCodeHandler{}).
		SetNext(&ResponseHandler{})

	chain.Handle(c)
	return
}
