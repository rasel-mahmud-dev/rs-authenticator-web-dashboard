package middlewares

import (
	"context"
	"net/http"
	context2 "rs/auth/app/context"
	"rs/auth/app/handlers/auth/verify"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		c := &context2.BaseContext{
			ResponseWriter: w,
			Request:        r,
		}

		chain := &verify.RequestValidationHandler{}
		chain.SetNext(&verify.ValidateAccessTokenHandler{})

		isNext := chain.Handle(c)

		ctx := context.WithValue(r.Context(), "authSession", c.AuthSession)
		r = r.WithContext(ctx)

		if isNext {
			next(w, r)
		}

	}
}
