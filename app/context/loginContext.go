package context

import (
	"rs/auth/app/dto"
	"rs/auth/app/models"
)

type LoginContext struct {
	LoginRequest dto.LoginRequest
	models.UserAuthAttempt
}
