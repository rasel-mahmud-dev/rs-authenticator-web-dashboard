package models

import "time"

type RecoveryCode struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Code      *string   `json:"code"`
	IsUsed    bool      `json:"is_used"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
