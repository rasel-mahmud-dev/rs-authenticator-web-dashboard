package login

import (
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	//
	//jsonHandler := &JSONValidationHandler{}
	//validationHandler := &RequestValidationHandler{}
	//existenceHandler := &UserExistenceHandler{}
	//passwordValidationHandler := &PasswordValidationHandler{}
	//generateJwtHandler := &GenerateJwtHandler{}
	//newSessionHandler := &authSession.NewSessionHandler{}
	//responseHandler := &ResponseHandler{}
	//
	//chain := jsonHandler
	//chain.SetNext(validationHandler).
	//	SetNext(existenceHandler).
	//	SetNext(passwordValidationHandler).
	//	SetNext(generateJwtHandler).
	//	SetNext(newSessionHandler).
	//	SetNext(responseHandler)
	//
	//chain.Handle(w, r)
}
