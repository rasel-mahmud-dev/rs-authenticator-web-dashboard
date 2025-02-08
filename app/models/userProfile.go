package models

import "time"

type UserProfile struct {
	UserID           string     `json:"user_id"`
	FullName         *string    `json:"full_name,omitempty"`
	AccountCreatedAt string     `json:"account_created_at,omitempty"`
	Avatar           *string    `json:"avatar,omitempty"`
	Cover            *string    `json:"cover,omitempty"`
	BirthDate        *time.Time `json:"birth_date,omitempty"`
	Gender           *string    `json:"gender,omitempty"`
	Phone            *string    `json:"phone,omitempty"`
	Location         *string    `json:"location,omitempty"`
	AboutMe          *string    `json:"about_me,omitempty"`
	Website          *string    `json:"website,omitempty"`

	// Social Media Links
	Facebook  *string `json:"facebook,omitempty"`
	Twitter   *string `json:"twitter,omitempty"`
	LinkedIn  *string `json:"linkedin,omitempty"`
	Instagram *string `json:"instagram,omitempty"`
	GitHub    *string `json:"github,omitempty"`
	YouTube   *string `json:"youtube,omitempty"`
	TikTok    *string `json:"tiktok,omitempty"`

	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
