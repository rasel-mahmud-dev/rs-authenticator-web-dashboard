package login

import (
	"net/http"
	"rs/auth/app/context"
	"rs/auth/app/handlers/auth/common"
	"rs/auth/app/handlers/authSession"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.BaseContext{
		ResponseWriter: w,
		Request:        r,
	}
	chain := &JSONValidationHandler{}
	chain.SetNext(&RequestValidationHandler{}).
		SetNext(&UserExistenceHandler{}).
		SetNext(&common.PasswordValidationHandler{}).
		SetNext(&GenerateJwtHandler{}).
		SetNext(&authSession.NewSessionHandler{}).
		SetNext(&ResponseHandler{})

	chain.Handle(ctx)
}
