package models

import "time"

type FileDownload struct {
	FileURL   string
	IP        string
	UserAgent string
	CreatedAt time.Time
}
