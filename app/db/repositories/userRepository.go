package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"rs/auth/app/cache"
	"rs/auth/app/db"
	"rs/auth/app/models"
	"rs/auth/app/utils"
	"sync"
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
	userC := cache.GetUserFromCache(email)
	if userC != nil {
		utils.LoggerInstance.Info("User from cache")
		return userC, nil
	}

	query := "SELECT id, name, email FROM users WHERE email = $1"
	var user models.User

	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		// added email to cache to store in nil
		cache.SetItem(email, &user)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no user found with email: %s", email)
		}
		return nil, fmt.Errorf("error l lll querying user by email: %w", err)
	}
	return &user, nil
}
