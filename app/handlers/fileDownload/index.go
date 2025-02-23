package fileDownload

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"rs/auth/app/models"
	"rs/auth/app/net/statusCode"
	fileDownloadRepository "rs/auth/app/repositories/fileDownload"
	"rs/auth/app/response"
	"time"
)

func FileDownload(w http.ResponseWriter, r *http.Request) {
	filePath := "./public/app-release.apk"

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Failed to open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Failed to get file info", http.StatusInternalServerError)
		return
	}

	ipAddress := r.RemoteAddr
	userAgent := r.UserAgent()

	fileDownloadRepository.FileDownloadInstance.Entry(models.FileDownload{
		FileURL:   filePath,
		IP:        ipAddress,
		UserAgent: userAgent,
		CreatedAt: time.Now(),
	})

	w.Header().Set("Content-Disposition", "attachment; filename=app-release.apk")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error streaming file", http.StatusInternalServerError)
		return
	}
}

func GetFileDownloadCount(w http.ResponseWriter, r *http.Request) {
	count := fileDownloadRepository.FileDownloadInstance.GetFileDownloadCount()
	response.Respond(w, statusCode.OK, "", map[string]int{
		"totalDownload": count,
	})
}
