package accountRecovery

import (
	"net/http"
	context2 "rs/auth/app/context"
	"rs/auth/app/handlers/auth/login"
	"rs/auth/app/handlers/authSession"
)

func AccountRecoveryChain(w http.ResponseWriter, r *http.Request) {

	c := &context2.BaseContext{
		RegistrationContext:  context2.RegistrationContext{},
		ResponseWriter:       w,
		Request:              r,
		Email:                "",
		User:                 nil,
		AuthSession:          nil,
		TwoFaSecurityContext: context2.TwoFaSecurityContext{},
	}

	chain := &PreparedContextState{}
	chain.SetNext(&ValidateRecoveryCodeHandler{}).
		SetNext(&PreparedContextForCreateNewSession{}).
		SetNext(&login.GenerateJwtHandler{}).
		SetNext(&authSession.NewSessionHandler{}).
		SetNext(&login.ResponseHandler{})

	chain.Handle(c)
	return
}
