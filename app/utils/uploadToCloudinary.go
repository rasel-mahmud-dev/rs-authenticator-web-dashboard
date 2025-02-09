package utils

import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"io"
	"mime/multipart"
	"os"
)

func UploadToCloudinary(file multipart.File, filename string) (string, error) {
	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return "", fmt.Errorf("failed to initialize Cloudinary: %v", err)
	}

	tempFile, err := os.CreateTemp("", "upload-*.jpg")
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %v", err)
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		return "", fmt.Errorf("failed to copy file: %v", err)
	}

	ctx := context.Background()
	uploadResult, err := cld.Upload.Upload(ctx, tempFile.Name(), uploader.UploadParams{
		PublicID:       filename,
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true),
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload to Cloudinary: %v", err)
	}

	return uploadResult.SecureURL, nil
}
