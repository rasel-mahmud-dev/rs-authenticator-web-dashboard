package repositories

import "rs/auth/internal/models"

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
}
