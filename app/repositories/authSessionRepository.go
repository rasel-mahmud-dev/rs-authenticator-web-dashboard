package repositories

import (
	"database/sql"
	"fmt"
	"rs/auth/app/db"
	"rs/auth/app/models"
	"rs/auth/app/utils"
	"time"
)

type authSessionRepository struct {
	db *sql.DB
}

var AuthSessionRepository *authSessionRepository

func init() {
	utils.LoggerInstance.Info("create user repo instance...")
	AuthSessionRepository = &authSessionRepository{db: db.GetDB()}
}

func (r *authSessionRepository) GetAuthSessionByAccessToken(token string) (*models.AuthSession, error) {
	//userC := cache.GetUserFromCache(email)
	//if userC != nil {
	//	utils.LoggerInstance.Info("User from cache")
	//	return userC, nil
	//}
	//
	//query := "SELECT id, username, password, email FROM users WHERE email = $1"
	var authSession models.AuthSession
	//
	//err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	//if err != nil {
	//	if errors.Is(err, sql.ErrNoRows) {
	//		return nil, nil
	//	}
	//	return nil, utils.Error("error l lll querying user by email: %w", err)
	//}
	//cache.SetItem(email, &user)
	return &authSession, nil
}

func (r *authSessionRepository) InsertAuthSession(payload models.AuthSession) (*models.AuthSession, error) {
	query := `
		INSERT INTO public.auth_sessions (
			user_id, ip_address, user_agent, access_token, refresh_token, 
			created_at, updated_at, last_used_at, is_revoked
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9
		) RETURNING id;
	`
	var newSession models.AuthSession
	currentTime := time.Now()
	err := r.db.QueryRow(
		query,
		payload.UserID,
		payload.IPAddress,
		payload.UserAgent,
		payload.AccessToken,
		payload.RefreshToken,
		currentTime,
		currentTime,
		nil,
		false,
	).
		Scan(&newSession.ID)

	if err != nil {
		return nil, fmt.Errorf("error creating auth session: %v", err)
	}
	newSession.IsRevoked = false
	return &newSession, nil
}
