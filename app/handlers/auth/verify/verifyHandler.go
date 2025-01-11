package verify

import (
	"net/http"
)

func AuthVerifyHandler(w http.ResponseWriter, r *http.Request) {

	validationHandler := &RequestValidationHandler{}
	validateAccessTokenHandler := &ValidateAccessTokenHandler{}
	responseHandler := &ResponseHandler{}

	chain := validationHandler
	chain.
		SetNext(validateAccessTokenHandler).
		SetNext(responseHandler)

	chain.Handle(w, &r)
}
