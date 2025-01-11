package repositories

import (
	"database/sql"
	"errors"
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
		false,      // is_active
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

func (r *mfaSecurityRepo) GetById(id string, userId string) (models.MfaSecurityToken, error) {
	query := `
		SELECT id, user_id, secret, recovery_codes, qr_code_url, is_active, created_at, updated_at, app_name, device_info
		FROM mfa_security_tokens
		WHERE id = $1 AND user_id = $2
	`
	var recoveryCodes []byte
	var token models.MfaSecurityToken

	err := r.db.QueryRow(query, id, userId).Scan(
		&token.ID,
		&token.UserID,
		&token.Secret,
		&recoveryCodes,
		&token.QrCodeURL,
		&token.IsActive,
		&token.CreatedAt,
		&token.UpdatedAt,
		&token.AppName,
		&token.DeviceInfo,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return token, fmt.Errorf("MFA security token not found")
		}
		return token, fmt.Errorf("failed to get MFA security token: %v", err)
	}

	if recoveryCodes != nil {
		err = pq.Array(&token.RecoveryCodes).Scan(recoveryCodes)
		if err != nil {
			return token, fmt.Errorf("failed to convert recovery codes: %v", err)
		}
	}
	return token, nil
}

func (r *mfaSecurityRepo) UpdateMfaSecurityToken(token models.MfaSecurityToken) error {
	query := `
		UPDATE mfa_security_tokens
		SET is_active = $1, updated_at = $2, app_name = $3
		WHERE user_id = $4 AND id = $5
	`
	_, err := r.db.Exec(query,
		token.IsActive,
		token.UpdatedAt,
		token.AppName,
		token.UserID,
		token.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update MFA security token: %v", err)
	}
	return nil
}
