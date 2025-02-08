package models

import (
	"time"
)

type UserAuthAttempt struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	AttemptType   string    `json:"attempt_type"`
	MFASecurityID string    `json:"mfa_security_id"`
	SecurityToken string    `json:"security_token"`
	IPAddress     string    `json:"ip_address"`
	UserAgent     string    `json:"user_agent"`
	LastAttemptAt time.Time `json:"last_attempt_at"`
	IsSuccessful  bool      `json:"is_successful"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
