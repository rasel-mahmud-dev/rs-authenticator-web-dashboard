package routes

import (
	"github.com/gorilla/mux"
	"rs/auth/app/handlers/fileDownload"
)

func FileDownload(router *mux.Router) {
	router.HandleFunc("/api/v1/download", fileDownload.FileDownload).Methods("GET")
	router.HandleFunc("/api/v1/download/count", fileDownload.GetFileDownloadCount).Methods("GET")
}
