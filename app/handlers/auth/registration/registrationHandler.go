package registration

import (
	"net/http"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {

	jsonHandler := &JSONValidationHandler{}
	validationHandler := &RequestValidationHandler{}

	checkExistenceUserHandler := &CheckExistenceUserHandler{}
	passwordValidationHandler := &PasswordValidationHandler{}
	createAccountHandler := &CreateAccountHandler{}
	authHandler := &AuthenticationHandler{}

	chain := jsonHandler
	chain.SetNext(validationHandler).
		SetNext(checkExistenceUserHandler).
		SetNext(passwordValidationHandler).
		SetNext(createAccountHandler).
		SetNext(authHandler)

	chain.Handle(w, &r)
}
