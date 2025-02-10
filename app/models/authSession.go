package models

import (
	"time"
)

type AuthSession struct {
	ID           string    `json:"id"`
	UserId       string    `json:"user_id"`
	IPAddress    string    `json:"ip_address"`
	UserAgent    string    `json:"user_agent"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	LastUsedAt   time.Time `json:"last_used_at"`
	IsRevoked    bool      `json:"is_revoked"`
	AuthMethod   string    `json:"auth_method,omitempty"`

	// populated field
	Avatar   string `json:"avatar,omitempty"`
	Cover    string `json:"cover,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}
