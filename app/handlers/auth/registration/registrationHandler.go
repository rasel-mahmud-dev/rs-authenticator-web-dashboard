package registration

import (
	"net/http"
	"rs/auth/app/context"
	"rs/auth/app/handlers/auth/common"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {

	baseContext := &context.BaseContext{
		ResponseWriter:      w,
		Request:             r,
		RegistrationContext: context.RegistrationContext{},
	}
	chain := &JSONValidationHandler{}
	chain.
		SetNext(&RequestValidationHandler{}).
		SetNext(&common.CheckExistenceUserHandler{}).
		SetNext(&PasswordValidationHandler{}).
		SetNext(&CreateAccountHandler{}).
		SetNext(&AuthenticationHandler{})

	chain.Handle(baseContext)
}
