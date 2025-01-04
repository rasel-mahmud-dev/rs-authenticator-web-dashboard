package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"rs/auth/internal/utils"
	"sync"

	"rs/auth/internal/db"
	"rs/auth/internal/models"
)

type userRepository struct {
	db *sql.DB
}

var (
	instance *userRepository
	once     sync.Once
)

func NewUserRepository() UserRepository {
	once.Do(func() {
		dbInstance := db.GetDB()
		utils.LoggerInstance.Info("create user repo instance...")
		instance = &userRepository{db: dbInstance}
	})

	return instance
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT id, name, email FROM users WHERE email = $1"
	var user models.User

	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no user found with email: %s", email)
		}
		return nil, fmt.Errorf("error querying user by email: %w", err)
	}
	return &user, nil
}
