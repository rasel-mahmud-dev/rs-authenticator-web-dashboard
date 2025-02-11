package trafficRepo

import (
	"database/sql"
	"log"
	"rs/auth/app/db"
	"rs/auth/app/dto"
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

func (r *Repository) GetTrafficDetailStats() ([]dto.DetailedTrafficStats, error) {
	query := `
		SELECT DATE(request_time) AS request_date, route_path, COUNT(*) AS request_count
		FROM user_traffic
		GROUP BY request_date, route_path
		ORDER BY request_date ASC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Println("Error fetching traffic stats:", err)
		return nil, err
	}
	defer rows.Close()

	var stats []dto.DetailedTrafficStats
	for rows.Next() {
		var s dto.DetailedTrafficStats
		err := rows.Scan(&s.RequestTime, &s.RoutePath, &s.Count)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		stats = append(stats, s)
	}

	return stats, nil
}
