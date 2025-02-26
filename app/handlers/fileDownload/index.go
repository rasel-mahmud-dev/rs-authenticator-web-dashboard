package fileDownload

import (
	"net/http"
	"rs/auth/app/models"
	"rs/auth/app/net/statusCode"
	fileDownloadRepository "rs/auth/app/repositories/fileDownload"
	"rs/auth/app/response"
	"time"
)

const googleDriveURL = "https://drive.google.com/file/d/1ldE3Xx-GvRO0iUykeIw3g_xlFY2JPWIJ/view?usp=sharing"

func FileDownload(w http.ResponseWriter, r *http.Request) {

	ipAddress := r.RemoteAddr
	userAgent := r.UserAgent()

	fileDownloadRepository.FileDownloadInstance.Entry(models.FileDownload{
		FileURL:   googleDriveURL,
		IP:        ipAddress,
		UserAgent: userAgent,
		CreatedAt: time.Now(),
	})

	response.Respond(w, statusCode.OK, "Success", map[string]string{
		"downloadURL": googleDriveURL,
	})
}

func GetFileDownloadCount(w http.ResponseWriter, r *http.Request) {
	count := fileDownloadRepository.FileDownloadInstance.GetFileDownloadCount()
	response.Respond(w, statusCode.OK, "", map[string]int{
		"totalDownload": count,
	})
}
