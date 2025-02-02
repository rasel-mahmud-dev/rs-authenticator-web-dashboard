package loginWithAuthenticator

import (
	"net/http"
	"rs/auth/app/context"
)

func LoginWithAuthenticator(w http.ResponseWriter, r *http.Request) {
	c := &context.BaseContext{
		RegistrationContext: context.RegistrationContext{},
		ResponseWriter:      w,
		Request:             r,
	}

	chain := &JSONValidationHandler{}
	chain.
		SetNext(&RequestValidationHandler{}).
		SetNext(&OtpVerificationHandler{})
	//SetNext(&login.GenerateJwtHandler{})
	//SetNext(&authSession.NewSessionHandler{}).
	//SetNext(&login.ResponseHandler{})

	chain.Handle(c)
}
