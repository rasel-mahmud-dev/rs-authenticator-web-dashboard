package generateRecoveryCode

import (
	"net/http"
	context2 "rs/auth/app/context"
)

func GenerateRecoveryCodeChain(w http.ResponseWriter, r *http.Request) {

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
	chain.SetNext(&AuthSessionHandler{}).
		SetNext(&CheckPreviousCodeHandler{}).
		SetNext(&GenerateRecoveryCodeHandler{}).
		SetNext(&ResponseHandler{})

	chain.Handle(c)
	return
}
