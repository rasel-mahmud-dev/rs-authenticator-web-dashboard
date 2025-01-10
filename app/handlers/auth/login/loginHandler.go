package login

import (
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	jsonHandler := &JSONValidationHandler{}
	validationHandler := &RequestValidationHandler{}

	existenceHandler := &UserExistenceHandler{}
	passwordValidationHandler := &PasswordValidationHandler{}
	authHandler := &AuthenticationHandler{}

	chain := jsonHandler
	chain.SetNext(validationHandler).
		SetNext(existenceHandler).
		SetNext(passwordValidationHandler).
		SetNext(authHandler)

	chain.Handle(w, r)
}
