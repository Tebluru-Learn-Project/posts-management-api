package service

import (
	"mime/multipart"
	"path/filepath"

	"go-api/internal/helper"
)

type UploadService struct{}

func NewUploadService() *UploadService {
	return &UploadService{}
}

func (s *UploadService) SaveImage(file *multipart.FileHeader, folder string) (string, string, error) {
	if err := helper.ValidateImage(file, 2*1024*1024); err != nil {
		return "", "", err
	}

	fileName := helper.GenerateFileName(file.Filename)
	uploadDir := filepath.Join("storage", "uploads", folder)
	dst := filepath.Join(uploadDir, fileName)

	if err := helper.EnsureDir(uploadDir); err != nil {
		return "", "", err
	}

	if err := helper.SaveMultipartFile(file, dst); err != nil {
		return "", "", err
	}

	return fileName, dst, nil
}