package generate2FASecret

import (
	"net/http"
)

func Generate2FASecretHandler(w http.ResponseWriter, r *http.Request) {

	authSessionHandler := &AuthSessionHandler{}
	checkInitTokenHandler := &CheckInitTokenHandler{}
	generateTotpSecretHandler := &GenerateTotpSecretHandler{}
	generateQRCodeHandler := &GenerateQRCodeHandler{}
	responseHandler := &ResponseHandler{}

	chain := authSessionHandler
	chain.SetNext(checkInitTokenHandler).
		SetNext(generateTotpSecretHandler).
		SetNext(generateQRCodeHandler).
		SetNext(responseHandler)

	chain.Handle(w, &r)
	return
}
