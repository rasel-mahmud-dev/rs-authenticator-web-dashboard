package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"rs/auth/app/cache"
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

func (r *authSessionRepository) GetAuthSessionByAccessToken(token string) *models.AuthSession {
	authSessionCached := cache.GetItem[*models.AuthSession](token)
	if authSessionCached != nil {
		utils.LoggerInstance.Info("auth session from cache")
		return authSessionCached
	}

	query := `SELECT 
    				s.id, 
       				s.is_revoked, 
				   s.access_token, 
				   s.refresh_token,
				   s.user_id, 
				   u.username, 
				   u.email,
				   COALESCE(u.avatar, '') AS avatar
FROM auth_sessions s 
				join public.users u 
					on u.id = s.user_id 
			WHERE access_token = $1`

	var authSession models.AuthSession

	err := r.db.QueryRow(query, token).Scan(
		&authSession.ID,
		&authSession.IsRevoked,
		&authSession.AccessToken,
		&authSession.RefreshToken,
		&authSession.UserId,
		&authSession.Username,
		&authSession.Email,
		&authSession.Avatar,
	)

	if err != nil {
		utils.LoggerInstance.Error(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return nil
	}
	return &authSession
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
		payload.UserId,
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
