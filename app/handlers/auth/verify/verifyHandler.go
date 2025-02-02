package verify

import (
	"net/http"
	context2 "rs/auth/app/context"
)

func AuthVerifyHandler(w http.ResponseWriter, r *http.Request) {

	c := context2.BaseContext{
		RegistrationContext:  context2.RegistrationContext{},
		AccessToken:          "",
		ResponseWriter:       w,
		Request:              r,
		Email:                "",
		User:                 nil,
		AuthSession:          nil,
		TwoFaSecurityContext: context2.TwoFaSecurityContext{},
	}

	chain := &RequestValidationHandler{}
	chain.
		SetNext(&ValidateAccessTokenHandler{}).
		SetNext(&ResponseHandler{})

	chain.Handle(c)
}
