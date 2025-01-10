package models

import (
	"time"
)

type AuthSessionHistory struct {
	ID           string    `json:"id"`
	UserId       string    `json:"user_id"`
	IPAddress    string    `json:"ip_address"`
	UserAgent    string    `json:"user_agent"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Status       string    `json:"status"`
	LastUsedAt   time.Time `json:"last_used_at"`
	IsRevoked    bool      `json:"is_revoked"`
	ArchivedAt   time.Time `json:"archived_at"`
}
