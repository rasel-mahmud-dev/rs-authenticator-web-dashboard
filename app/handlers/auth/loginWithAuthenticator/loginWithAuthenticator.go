package loginWithAuthenticator

import (
	"net/http"
)

func LoginWithAuthenticator(w http.ResponseWriter, r *http.Request) {
	jsonHandler := &JSONValidationHandler{}
	validationHandler := &RequestValidationHandler{}
	otpVerificationHandler := &OtpVerificationHandler{}

	chain := jsonHandler
	chain.
		SetNext(validationHandler).
		SetNext(otpVerificationHandler)

	chain.Handle(w, &r)
}
