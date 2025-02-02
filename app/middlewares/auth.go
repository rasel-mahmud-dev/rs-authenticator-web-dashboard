package middlewares

import (
	"net/http"
	"rs/auth/app/context"
	"rs/auth/app/handlers/auth/verify"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		c := context.BaseContext{
			ResponseWriter: w,
			Request:        r,
		}

		validationHandler := &verify.RequestValidationHandler{}
		validateAccessTokenHandler := &verify.ValidateAccessTokenHandler{}

		chain := validationHandler
		chain.SetNext(validateAccessTokenHandler)

		chain.Handle(c)
		next(w, r)
	}
}
