package routes

import (
	"github.com/gorilla/mux"
	"rs/auth/app/handlers/mfaSecurity"
	"rs/auth/app/handlers/mfaSecurity/generate2FASecret"
	"rs/auth/app/handlers/mfaSecurity/generateRecoveryCode"
	"rs/auth/app/middlewares"
)

func MultiFactorAuthenticationRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/mfa/generate-2fa-secret", middlewares.Auth(generate2FASecret.Generate2FASecretHandler)).Methods("POST")
	router.HandleFunc("/api/v1/mfa/generate-recovery-code", middlewares.Auth(generateRecoveryCode.GenerateRecoveryCodeChain)).Methods("POST")
	router.HandleFunc("/api/v1/mfa/generate-2fa-secret-complete", middlewares.Auth(mfaSecurity.Finalize2FASecret)).Methods("POST")
	router.HandleFunc("/api/v1/mfa/unlink-authenticator/{id}", middlewares.Auth(mfaSecurity.RemoveAuthenticator)).Methods("POST")
	router.HandleFunc("/api/v1/mfa/authenticated-apps", middlewares.Auth(mfaSecurity.GetAllConnectedAuthenticatorApps)).Methods("GET")
}
