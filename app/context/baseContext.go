package context

import (
	"net/http"
	"rs/auth/app/models"
)

type BaseContext struct {
	RegistrationContext
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	AccessToken    string
	Email          string
	User           *models.User
	AuthSession    *models.AuthSession
	TwoFaSecurityContext
	LoginContext
}
