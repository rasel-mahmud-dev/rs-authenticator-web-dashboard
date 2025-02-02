package login

import (
	"net/http"
	"rs/auth/app/handlers/auth/common"
	"rs/auth/app/handlers/authSession"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	chain := &JSONValidationHandler{}
	chain.SetNext(&RequestValidationHandler{}).
		SetNext(&UserExistenceHandler{}).
		SetNext(&authCommon.PasswordValidationHandler{}).
		SetNext(&GenerateJwtHandler{}).
		SetNext(&authSession.NewSessionHandler{}).
		SetNext(&ResponseHandler{})

	chain.Handle(w, &r)
}
