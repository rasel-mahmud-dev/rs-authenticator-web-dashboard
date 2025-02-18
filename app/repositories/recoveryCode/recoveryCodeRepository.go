package recoveryCode

import (
	"database/sql"
	"errors"
	"fmt"
	"rs/auth/app/db"
	"rs/auth/app/models"
	"rs/auth/app/utils"
	"strings"
)

type Repository struct {
	db *sql.DB
}

var RecoveryCodeRepository *Repository

func init() {
	utils.LoggerInstance.Info("create user repo instance...")
	RecoveryCodeRepository = &Repository{db: db.GetDB()}
}

//func (r *Repository) InsertRecoveryCodes(traffic models.RecoveryCode) error {
//	query := `
//		INSERT INTO user_traffic (
//			route_path, http_method, user_agent, ip_address, request_time, response_time
//		) VALUES ($1, $2, $3, $4, $5, $6)
//	`
//
//	_, err := r.db.Exec(query,
//		traffic.RoutePath,
//		traffic.HTTPMethod,
//		traffic.UserAgent,
//		traffic.IPAddress,
//		traffic.RequestTime,
//		traffic.ResponseTime,
//	)
//
//	return err
//}

func (r *Repository) InsertMultipleRecoveryCodes(recoveryCodes []models.RecoveryCode) error {
	if len(recoveryCodes) == 0 {
		return nil // No data to insert
	}

	query := `INSERT INTO recovery_codes (user_id, code, is_used, created_at, updated_at, expires_at) VALUES`
	values := []interface{}{}
	placeholders := []string{}

	for i, recovery := range recoveryCodes {
		n := i * 6
		placeholders = append(placeholders,
			fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)",
				n+1, n+2, n+3, n+4, n+5, n+6))

		values = append(values, recovery.UserID, recovery.Code,
			recovery.IsUsed, recovery.CreatedAt, recovery.UpdatedAt, recovery.ExpiresAt)
	}

	finalQuery := query + strings.Join(placeholders, ",")

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(finalQuery, values...)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *Repository) GetLast10RecoveryCodes(userID string) ([]models.RecoveryCode, error) {
	query := `
		SELECT id, user_id, code, is_used, created_at, updated_at, expires_at
		FROM recovery_codes
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT 10
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recoveryCodes []models.RecoveryCode
	for rows.Next() {
		var recovery models.RecoveryCode
		err := rows.Scan(&recovery.ID, &recovery.UserID, &recovery.Code, &recovery.IsUsed, &recovery.CreatedAt, &recovery.UpdatedAt, &recovery.ExpiresAt)
		if err != nil {
			return nil, err
		}
		recoveryCodes = append(recoveryCodes, recovery)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return recoveryCodes, nil
}

func (r *Repository) GetValidRecoveryCode(code string) (models.RecoveryCode, error) {
	query := `
		SELECT id, user_id, code, is_used, created_at, updated_at, expires_at
		FROM recovery_codes
		WHERE code = $1
		AND is_used = FALSE
		AND expires_at > NOW()
		ORDER BY created_at DESC
		LIMIT 1
	`

	var recovery models.RecoveryCode
	err := r.db.QueryRow(query, code).Scan(&recovery.ID, &recovery.UserID, &recovery.Code, &recovery.IsUsed, &recovery.CreatedAt, &recovery.UpdatedAt, &recovery.ExpiresAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.RecoveryCode{}, nil
		}
		return models.RecoveryCode{}, err
	}

	return recovery, nil
}

func (r *Repository) MakeInvalidRecoveryCodeById(id string) error {
	query := `
		UPDATE recovery_codes SET expires_at = NOW(), is_used = TRUE
		WHERE id = $1
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
