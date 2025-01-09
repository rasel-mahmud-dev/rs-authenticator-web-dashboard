package repositories

import (
	"rs/auth/app/models"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
}
