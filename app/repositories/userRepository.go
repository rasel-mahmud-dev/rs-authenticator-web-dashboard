package repositories

import (
	"database/sql"
	"errors"
	"rs/auth/app/cache"
	"time"

	"rs/auth/app/db"
	"rs/auth/app/models"
	"rs/auth/app/utils"
	"sync"
)

type UserRepository struct {
	db *sql.DB
}

var (
	instance *UserRepository
	once     sync.Once
)

func NewUserRepository() *UserRepository {
	once.Do(func() {
		dbInstance := db.GetDB()
		utils.LoggerInstance.Info("create user repo instance...")
		instance = &UserRepository{db: dbInstance}
	})
	return instance
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	userC := cache.GetUserFromCache(email)
	if userC != nil {
		utils.LoggerInstance.Info("User from cache")
		return userC, nil
	}

	query := "SELECT id, username, password, email FROM users WHERE email = $1"
	var user models.User

	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, utils.Error("error l lll querying user by email: %w", err)
	}
	cache.SetItem(email, &user)
	return &user, nil
}

func (r *UserRepository) CreateAccount(user models.User) (*models.User, error) {
	query := `
		INSERT INTO public.users (username, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, username, email, password, created_at, updated_at`

	var newUser models.User

	createUpdateTime := time.Now()

	err := r.db.QueryRow(
		query,
		user.Username,
		user.Email,
		user.Password,
		createUpdateTime,
		createUpdateTime,
	).
		Scan(
			&newUser.ID,
			&newUser.Username,
			&newUser.Email,
			&newUser.Password,
			&newUser.CreatedAt,
			&newUser.UpdatedAt,
		)

	if err != nil {
		return nil, utils.Error("error creating user account: %w", err)
	}
	return &newUser, nil
}
