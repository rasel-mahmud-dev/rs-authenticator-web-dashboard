package models

import "time"

type MfaSecurityToken struct {
	ID            string     `json:"id"`
	UserID        string     `json:"user_id"`
	CodeName      *string    `json:"code_name"`
	Secret        string     `json:"secret"`
	RecoveryCodes []string   `json:"recovery_codes"`
	QrCodeURL     *string    `json:"qr_code_url"`
	IsActive      bool       `json:"is_active"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	LinkedAt      *time.Time `json:"linked_at"`
	AppName       string     `json:"app_name"`
	DeviceInfo    *string    `json:"device_info"`
	IsInit        bool       `json:"is_init"`
}
