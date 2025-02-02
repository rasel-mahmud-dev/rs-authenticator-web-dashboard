package registration

import (
	"net/http"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {

	chain := &JSONValidationHandler{}
	chain.
		SetNext(&RequestValidationHandler{}).
		SetNext(&CheckExistenceUserHandler{}).
		SetNext(&PasswordValidationHandler{}).
		SetNext(&CreateAccountHandler{}).
		SetNext(&AuthenticationHandler{})

	chain.Handle(w, &r)
}
