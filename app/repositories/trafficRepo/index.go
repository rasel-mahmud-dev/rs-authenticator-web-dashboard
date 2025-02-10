package trafficRepo

import (
	"database/sql"
	"rs/auth/app/db"
	"rs/auth/app/models"
	"rs/auth/app/utils"
)

type Repository struct {
	db *sql.DB
}

var TrafficRepository *Repository

func init() {
	utils.LoggerInstance.Info("create user repo instance...")
	TrafficRepository = &Repository{db: db.GetDB()}
}

func (r *Repository) InsertApiTraffic(traffic models.UserTraffic) error {
	query := `
		INSERT INTO user_traffic (
			route_path, http_method, user_agent, ip_address, request_time, response_time
		) VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(query,
		traffic.RoutePath,
		traffic.HTTPMethod,
		traffic.UserAgent,
		traffic.IPAddress,
		traffic.RequestTime,
		traffic.ResponseTime,
	)

	return err
}
