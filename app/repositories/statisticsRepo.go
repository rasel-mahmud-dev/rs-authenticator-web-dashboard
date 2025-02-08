package repositories

import (
	"rs/auth/app/dto"
	"rs/auth/app/utils"
	"time"
)

func (r *UserRepository) GetUserRegistrationStats() ([]dto.UserRegistrationStats, error) {
	now := time.Now()
	startDate := now.AddDate(0, 0, -30) // 30 days ago
	endDate := now                      // Today

	query := `SELECT 
			DATE(created_at) AS date, 
			COUNT(*) AS count
		FROM users 
		WHERE created_at BETWEEN $1 AND $2
		GROUP BY DATE(created_at)
		ORDER BY DATE(created_at);
	`

	rows, err := r.db.Query(query, startDate, endDate)
	if err != nil {
		return nil, utils.Error("Error querying user registrations: %w", err)
	}
	defer rows.Close()

	var stats []dto.UserRegistrationStats

	for rows.Next() {
		var stat dto.UserRegistrationStats
		if err := rows.Scan(&stat.Date, &stat.Count); err != nil {
			return nil, utils.Error("Error scanning user registration stats: %w", err)
		}
		stats = append(stats, stat)
	}

	if err := rows.Err(); err != nil {
		return nil, utils.Error("Error iterating over rows: %w", err)
	}
	return stats, nil
}

func (r *UserRepository) GetAuthenticationStats() ([]dto.AuthenticatorStats, error) {
	now := time.Now()
	startDate := now.AddDate(0, 0, -30) // Last 30 days

	query := `
		SELECT 
			DATE(created_at) AS date,
			COUNT(CASE WHEN auth_method = 'authenticator' THEN 1 END) AS authenticator,
			COUNT(CASE WHEN auth_method = 'password' THEN 1 END) AS password
		FROM auth_sessions
		WHERE created_at >= $1
		GROUP BY DATE(created_at)
		ORDER BY DATE(created_at) ASC;
	`

	rows, err := r.db.Query(query, startDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []dto.AuthenticatorStats
	for rows.Next() {
		var stat dto.AuthenticatorStats
		err := rows.Scan(&stat.Date, &stat.Authenticator, &stat.Password)
		if err != nil {
			return nil, err
		}
		stats = append(stats, stat)
	}

	return stats, nil
}

func (r *UserRepository) GetAttemptRateStats() dto.GetAttemptRateStatsResult {
	query := `
		SELECT 
			(SELECT count(id) FROM user_auth_attempts) AS failed,
			(SELECT count(id) FROM auth_sessions) AS success;
	`

	var result dto.GetAttemptRateStatsResult
	err := r.db.QueryRow(query).Scan(&result.Failed, &result.Success)
	if err != nil {
		utils.LoggerInstance.Error(err.Error())
	}

	result.Total = result.Failed + result.Success
	return result
}

func (r *UserRepository) GetAttemptRateDetailStats() ([]dto.AttemptRateStatsDetail, error) {
	failedQuery := `
        SELECT 
            DATE(created_at) AS date, 
            COUNT(id) AS failed
        FROM user_auth_attempts
        GROUP BY DATE(created_at)
        ORDER BY date;
    `
	rows, err := r.db.Query(failedQuery)
	if err != nil {
		utils.LoggerInstance.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	failedCounts := make(map[string]int)

	for rows.Next() {
		var date time.Time
		var failed int
		if err := rows.Scan(&date, &failed); err != nil {
			utils.LoggerInstance.Error(err.Error())
			return nil, err
		}
		failedCounts[date.Format("2006-01-02")] = failed
	}

	// Query for successful attempts
	successQuery := `
        SELECT 
            DATE(created_at) AS date, 
            COUNT(id) AS success
        FROM auth_sessions
        GROUP BY DATE(created_at)
        ORDER BY date;
    `
	rows, err = r.db.Query(successQuery)
	if err != nil {
		utils.LoggerInstance.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	var result []dto.AttemptRateStatsDetail

	for rows.Next() {
		var date time.Time
		var success int
		if err := rows.Scan(&date, &success); err != nil {
			utils.LoggerInstance.Error(err.Error())
			return nil, err
		}

		detail := dto.AttemptRateStatsDetail{
			Date:    date,
			Success: success,
		}

		if failed, exists := failedCounts[date.Format("2006-01-02")]; exists {
			detail.Failed = failed
		} else {
			detail.Failed = 0
		}

		result = append(result, detail)
	}

	return result, nil
}
