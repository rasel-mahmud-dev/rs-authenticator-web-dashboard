package userProfile

import (
	"database/sql"
	"errors"
	"fmt"
	"rs/auth/app/db"
	"rs/auth/app/dto"
	"rs/auth/app/models"
	"rs/auth/app/utils"
	"sync"
	"time"
)

type Repository struct {
	db *sql.DB
}

var (
	instance *Repository
	once     sync.Once
)

func NewRepository() *Repository {
	once.Do(func() {
		dbInstance := db.GetDB()
		utils.LoggerInstance.Info("create user repo instance...")
		instance = &Repository{db: dbInstance}
	})
	return instance
}

func (r *Repository) InsertOrUpdateUserProfile(profile dto.UpdateProfilePayload) error {

	query := `
		INSERT INTO user_profiles (
			user_id, full_name, birth_date, gender, phone, location, about_me, website,
			facebook, twitter, linkedin, instagram, github, youtube, tiktok, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, NOW(), NOW())
		ON CONFLICT (user_id) 
		DO UPDATE SET 
			full_name = EXCLUDED.full_name,
			birth_date = EXCLUDED.birth_date,
			gender = EXCLUDED.gender,
			phone = EXCLUDED.phone,
			location = EXCLUDED.location,
			about_me = EXCLUDED.about_me,
			website = EXCLUDED.website,
			facebook = EXCLUDED.facebook,
			twitter = EXCLUDED.twitter,
			linkedin = EXCLUDED.linkedin,
			instagram = EXCLUDED.instagram,
			github = EXCLUDED.github,
			youtube = EXCLUDED.youtube,
			tiktok = EXCLUDED.tiktok,
			updated_at = NOW();
	`

	parsedTime, err := time.Parse("2006-01-02", *profile.BirthDate)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}

	fmt.Println("parsedTime", parsedTime)

	_, err = r.db.Exec(query,
		profile.UserID, profile.FullName, parsedTime, profile.Gender,
		profile.Phone, profile.Location, profile.AboutMe, profile.Website,
		profile.Facebook, profile.Twitter, profile.LinkedIn, profile.Instagram,
		profile.GitHub, profile.YouTube, profile.TikTok,
	)

	return err
}

func (r *Repository) GetUserProfile(userID string) (*models.UserProfile, error) {
	query := `
		SELECT 
			u.id, u.avatar, u.created_at AS account_created_at, 
			COALESCE(p.full_name, ''), p.birth_date, p.gender, p.phone, p.location, p.about_me, p.website,
			p.facebook, p.twitter, p.linkedin, p.instagram, p.github, p.youtube, p.tiktok, 
			p.created_at, p.updated_at
		FROM users u
		LEFT JOIN user_profiles p ON u.id = p.user_id
		WHERE u.id = $1;
	`

	var profile models.UserProfile
	var createdAt, updatedAt sql.NullTime // Use sql.NullTime to handle NULL timestamps

	err := r.db.QueryRow(query, userID).Scan(
		&profile.UserID, &profile.Avatar, &profile.AccountCreatedAt,
		&profile.FullName, &profile.BirthDate, &profile.Gender,
		&profile.Phone, &profile.Location, &profile.AboutMe, &profile.Website,
		&profile.Facebook, &profile.Twitter, &profile.LinkedIn, &profile.Instagram,
		&profile.GitHub, &profile.YouTube, &profile.TikTok,
		&createdAt, &updatedAt, // Use sql.NullTime
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No profile found
		}
		return nil, err
	}

	if createdAt.Valid {
		profile.CreatedAt = &createdAt.Time
	} else {
		profile.CreatedAt = nil
	}

	if updatedAt.Valid {
		profile.UpdatedAt = &updatedAt.Time
	} else {
		profile.UpdatedAt = nil
	}

	return &profile, nil
}
