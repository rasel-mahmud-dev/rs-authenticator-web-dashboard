package login

import (
	"net/http"
	"rs/auth/app/db/repositories"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	jsonHandler := &JSONValidationHandler{}
	validationHandler := &RequestValidationHandler{}

	userRepo := repositories.NewUserRepository()
	existenceHandler := NewUserExistenceHandler(userRepo)
	authHandler := &AuthenticationHandler{}

	jsonHandler.SetNext(validationHandler)
	validationHandler.SetNext(existenceHandler)
	existenceHandler.SetNext(authHandler)

	jsonHandler.Handle(w, r)
}
