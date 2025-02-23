package fileDownloadRepository

import (
	"database/sql"
	"rs/auth/app/db"
	"rs/auth/app/models"
	"rs/auth/app/utils"
)

type FileDownload struct {
	db *sql.DB
}

var FileDownloadInstance *FileDownload

func init() {
	utils.LoggerInstance.Info("create user repo instance...")
	FileDownloadInstance = &FileDownload{db: db.GetDB()}
}

func (r *FileDownload) GetFileDownloadCount() int {
	query := `SELECT count(id) as total_download from file_downloads`
	var totalDownload int
	err := r.db.QueryRow(query).Scan(&totalDownload)
	if err != nil {
		utils.LoggerInstance.Error("failed to retrieve download count: %w", err)
		return 0
	}
	return totalDownload
}

func (r *FileDownload) Entry(item models.FileDownload) int {
	query := `
		INSERT INTO file_downloads (file_url, ip_address, user_agent, created_at)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(query, item.FileURL, item.IP, item.UserAgent, item.CreatedAt)
	if err != nil {
		utils.LoggerInstance.Error("failed to store download count: %w", err)
		return 0
	}

	return 0
}
