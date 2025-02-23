package repositories

import (
	"database/sql"
	"errors"
	"rs/auth/app/cache"
	"rs/auth/app/db"
	"rs/auth/app/models"
	"rs/auth/app/utils"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

var UserRepositoryInstance *UserRepository

func init() {
	utils.LoggerInstance.Info("create user repo instance...")
	UserRepositoryInstance = &UserRepository{db: db.GetDB()}
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	userC := cache.GetUserFromCache(email)
	if userC != nil {
		utils.LoggerInstance.Info("User from cache")
		return userC, nil
	}

	query := `SELECT 
    u.id, 
    u.username, 
    u.password, 
    u.email, 
    COALESCE(p.avatar, '') AS avatar FROM users u
		LEFT JOIN public.user_profiles p 
		    ON u.id = p.user_id 
 	WHERE email = $1`
	var user models.User

	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Avatar)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, utils.Error("error l lll querying user by email: %w", err)
	}
	cache.SetItem(email, &user)
	return &user, nil
}

func (r *UserRepository) GetAllUsers(page int, limit int) ([]models.User, int, error) {
	offset := (page - 1) * limit

	query := `
		SELECT u.id, u.username, u.email, u.created_at, COALESCE(up.avatar, '') AS avatar 
		FROM users u left join public.user_profiles up on u.id = up.user_id
		ORDER BY u.created_at DESC 
		LIMIT $1 OFFSET $2
	`

	var users []models.User

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, utils.Error("error querying users: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.Avatar)
		if err != nil {
			return nil, 0, utils.Error("error scanning user data: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, utils.Error("error iterating over rows: %w", err)
	}

	countQuery := `SELECT COUNT(id) FROM users`
	var totalItems int

	err = r.db.QueryRow(countQuery).Scan(&totalItems)
	if err != nil {

	}

	return users, totalItems, nil
}

func (r *UserRepository) GetUserById(userId string) (*models.User, error) {
	userC := cache.GetUserFromCache(userId)
	if userC != nil {
		utils.LoggerInstance.Info("User from cache")
		return userC, nil
	}

	query := "SELECT id, username, password, email, COALESCE(avatar, '') AS avatar FROM users WHERE id = $1"
	var user models.User

	err := r.db.QueryRow(query, userId).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Avatar)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, utils.Error("error l lll querying user by email: %w", err)
	}
	cache.SetItem(userId, &user)
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
