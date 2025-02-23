package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
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
		INSERT INTO mfa_security_tokens (user_id, secret, qr_code_url, is_active, created_at, updated_at, app_name,  code_name)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, user_id, secret, qr_code_url, is_active, created_at, updated_at, app_name, code_name
	`

	var savedToken models.MfaSecurityToken
	err := r.db.QueryRow(query,
		token.UserID,
		token.Secret,
		&token.QrCodeURL,
		&token.IsActive, // is_active
		time.Now(),      // created_at
		time.Now(),      // updated_at
		token.AppName,
		token.CodeName).Scan(
		&savedToken.ID,
		&savedToken.UserID,
		&savedToken.Secret,
		&savedToken.QrCodeURL,
		&savedToken.IsActive,
		&savedToken.CreatedAt,
		&savedToken.UpdatedAt,
		&savedToken.AppName,
		&savedToken.CodeName,
	)

	if err != nil {
		fmt.Println(err)
		return savedToken, fmt.Errorf("failed to insert MFA security token: %v", err)
	}

	return savedToken, nil
}

func (r *mfaSecurityRepo) GetById(id string, userId string) (models.MfaSecurityToken, error) {
	query := `
		SELECT id, user_id, secret,  qr_code_url, is_active, created_at, updated_at, app_name,  code_name
		FROM mfa_security_tokens
		WHERE id = $1 AND user_id = $2
	`
	var token models.MfaSecurityToken

	err := r.db.QueryRow(query, id, userId).Scan(
		&token.ID,
		&token.UserID,
		&token.Secret,
		&token.QrCodeURL,
		&token.IsActive,
		&token.CreatedAt,
		&token.UpdatedAt,
		&token.AppName,
		&token.CodeName,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return token, fmt.Errorf("MFA security token not found")
		}
		return token, fmt.Errorf("failed to get MFA security token: %v", err)
	}

	return token, nil
}

func (r *mfaSecurityRepo) GetLastInit(userId string) (*models.MfaSecurityToken, error) {
	query := `
		SELECT id, user_id, secret, qr_code_url, is_active, created_at, updated_at, app_name,  code_name
		FROM mfa_security_tokens
		WHERE user_id = $1 limit 1
	`
	var token models.MfaSecurityToken

	err := r.db.QueryRow(query, userId).Scan(
		&token.ID,
		&token.UserID,
		&token.Secret,
		&token.QrCodeURL,
		&token.IsActive,
		&token.CreatedAt,
		&token.UpdatedAt,
		&token.AppName,
		&token.CodeName,
	)

	if err != nil {
		utils.LoggerInstance.Debug(err.Error())
		return nil, fmt.Errorf("failed to get MFA security token: %v", err)
	}

	return &token, nil
}

func (r *mfaSecurityRepo) ResetInitToken(userId string) error {
	query := `UPDATE mfa_security_tokens 
		set is_active = false
    WHERE user_id = $1`

	_, err := r.db.Exec(query, userId)
	if err != nil {
		utils.LoggerInstance.Debug(err.Error())
		return fmt.Errorf("failed to get MFA security token: %v", err)
	}
	return nil
}

func (r *mfaSecurityRepo) GetAllItems(userId string) ([]models.MfaSecurityToken, error) {
	query := `
		SELECT id, user_id, secret, qr_code_url, is_active, created_at, updated_at, app_name, code_name, linked_at
		FROM mfa_security_tokens where user_id = $1 AND is_active = true
	`

	var tokens []models.MfaSecurityToken
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch MFA security tokens: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var token models.MfaSecurityToken
		if err := rows.Scan(
			&token.ID,
			&token.UserID,
			&token.Secret,
			&token.QrCodeURL,
			&token.IsActive,
			&token.CreatedAt,
			&token.UpdatedAt,
			&token.AppName,
			&token.CodeName,
			&token.LinkedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		tokens = append(tokens, token)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while fetching rows: %v", err)
	}
	return tokens, nil
}

func (r *mfaSecurityRepo) UpdateMfaSecurityToken(token models.MfaSecurityToken) error {
	query := `
		UPDATE mfa_security_tokens
		SET is_active = $1, updated_at = $2, app_name = $3, linked_at = $2
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

func (r *mfaSecurityRepo) RemoveAuthenticator(userId string, id string) error {
	query := `
		UPDATE mfa_security_tokens
		SET is_active = false, updated_at = CURRENT_TIMESTAMP
		WHERE user_id = $1 AND id = $2 
	`
	_, err := r.db.Exec(query, userId, id)
	if err != nil {
		return fmt.Errorf("failed to disable MFA security token: %v", err)
	}
	return nil
}

func validateOtpCode(otpCode string, secret string) (bool, error) {
	opts := totp.ValidateOpts{
		Period:    30,
		Skew:      0,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA256,
	}
	return totp.ValidateCustom(otpCode, secret, time.Now().UTC(), opts)
}

func (r *mfaSecurityRepo) VerifyMfaPasscode(userId string, otpCode string) (string, error) {
	query := `SELECT user_id, secret FROM mfa_security_tokens WHERE user_id = $1 AND is_active = TRUE`

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return "", err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var userID string
	var secret string

	for rows.Next() {
		err := rows.Scan(&userID, &secret)
		if err != nil {
			return "", fmt.Errorf("failed to read MFA token row: %v", err)
		}

		isMatch, err := validateOtpCode(otpCode, secret)
		if isMatch {
			return userID, nil
		}
	}
	return "", errors.New("invalid otp code")
}

func (r *mfaSecurityRepo) Is2FaEnabled(userId string) bool {
	query := `SELECT count(user_id) as count FROM mfa_security_tokens WHERE user_id = $1 AND is_active = TRUE`

	var connectedItem int
	err := r.db.QueryRow(query, userId).Scan(&connectedItem)
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return false
	}
	return connectedItem > 0
}
