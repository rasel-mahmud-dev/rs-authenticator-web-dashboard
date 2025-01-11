package repositories

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"rs/auth/app/db"
	"rs/auth/app/models"
	"rs/auth/app/utils"
	"time"
)

type mfaSecurityRepo struct {
	db *sql.DB
}

var MfaSecurityTokenRepo *mfaSecurityRepo

func init() {
	utils.LoggerInstance.Info("create MfaSecurityTokenRepo instance...")
	MfaSecurityTokenRepo = &mfaSecurityRepo{db: db.GetDB()}
}
func (r *mfaSecurityRepo) InsertMfaSecurityToken(token models.MfaSecurityToken) (models.MfaSecurityToken, error) {
	query := `
		INSERT INTO mfa_security_tokens (user_id, secret, recovery_codes, qr_code_url, is_active, created_at, updated_at, app_name, device_info)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, user_id, secret, recovery_codes, qr_code_url, is_active, created_at, updated_at, app_name, device_info
	`

	//marshal, err := json.Marshal(token.RecoveryCodes)
	//if err != nil {
	//	return models.MfaSecurityToken{}, err
	//}
	var recoveryCodes []byte

	var savedToken models.MfaSecurityToken
	err := r.db.QueryRow(query,
		token.UserID,
		token.Secret,
		pq.Array(token.RecoveryCodes),
		&token.QrCodeURL,
		true,       // is_active
		time.Now(), // created_at
		time.Now(), // updated_at
		token.AppName,
		token.DeviceInfo,
	).Scan(
		&savedToken.ID,
		&savedToken.UserID,
		&savedToken.Secret,
		&recoveryCodes,
		&savedToken.QrCodeURL,
		&savedToken.IsActive,
		&savedToken.CreatedAt,
		&savedToken.UpdatedAt,
		&savedToken.AppName,
		&savedToken.DeviceInfo,
	)

	if err != nil {
		fmt.Println(err)
		return savedToken, fmt.Errorf("failed to insert MFA security token: %v", err)
	}

	if recoveryCodes != nil {
		err = pq.Array(&savedToken.RecoveryCodes).Scan(recoveryCodes)
		if err != nil {
			utils.LoggerInstance.Warn("failed to convert recovery codes: %v", err)
		}
	}

	return savedToken, nil
}
