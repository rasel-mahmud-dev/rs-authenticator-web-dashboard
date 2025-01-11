package repositories

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"rs/auth/app/cache"
	"rs/auth/app/db"
	"rs/auth/app/models"
	"rs/auth/app/utils"
	"time"
)

type twoFactorAuthRepo struct {
	db *sql.DB
}

var TwoFactorAuthRepo *twoFactorAuthRepo

func init() {
	utils.LoggerInstance.Info("create user repo instance...")
	TwoFactorAuthRepo = &twoFactorAuthRepo{db: db.GetDB()}
}

func (r *twoFactorAuthRepo) Save2FASecret(userID string, secret string, qrCodeData []byte) error {
	query := `
        INSERT INTO user_security_tokens(user_id, secret, qr_code_url, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
    `
	// Convert QR code data to base64 for storage.
	qrCodeBase64 := fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(qrCodeData))

	_, err := r.db.Exec(query, userID, secret, qrCodeBase64, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (r *twoFactorAuthRepo) GetAuthSessionByAccessToken(token string) *models.UserSecurityToken {
	authSessionCached := cache.GetItem[*models.UserSecurityToken](token)
	if authSessionCached != nil {
		utils.LoggerInstance.Info("auth session from cache")
		return authSessionCached
	}

	return authSessionCached
}
