package middlewares

import (
	"net/http"
	"rs/auth/app/handlers/auth/verify"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		validationHandler := &verify.RequestValidationHandler{}
		validateAccessTokenHandler := &verify.ValidateAccessTokenHandler{}

		chain := validationHandler
		chain.SetNext(validateAccessTokenHandler)

		chain.Handle(w, &r)
		next(w, r)
	}
}
