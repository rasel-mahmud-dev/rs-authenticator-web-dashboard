package loginWithAuthenticator

import (
	"net/http"
	"rs/auth/app/handlers/auth/login"
	"rs/auth/app/handlers/authSession"
)

func LoginWithAuthenticator(w http.ResponseWriter, r *http.Request) {
	jsonHandler := &JSONValidationHandler{}
	validationHandler := &RequestValidationHandler{}
	otpVerificationHandler := &OtpVerificationHandler{}
	generateJwtHandler := &login.GenerateJwtHandler{}
	newSessionHandler := &authSession.NewSessionHandler{}
	responseHandler := &login.ResponseHandler{}

	chain := jsonHandler
	chain.
		SetNext(validationHandler).
		SetNext(otpVerificationHandler).
		SetNext(generateJwtHandler).
		SetNext(newSessionHandler).
		SetNext(responseHandler)

	chain.Handle(w, &r)
}
