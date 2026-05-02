package helper

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

var allowedImageMimeTypes = map[string]bool{
	"image/jpeg": true,
	"image/jpg":  true,
	"image/png":  true,
	"image/webp": true,
}

func ValidateImage(file *multipart.FileHeader, maxSize int64) error {
	if file.Size > maxSize {
		return errors.New("file size exceeds limit")
	}

	contentType := file.Header.Get("Content-Type")
	if !allowedImageMimeTypes[contentType] {
		return errors.New("invalid file type")
	}

	return nil
}

func GenerateFileName(originalName string) string {
	ext := strings.ToLower(filepath.Ext(originalName))
	return uuid.New().String() + ext
}

func SaveMultipartFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func EnsureDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func ResolveFileURL(path string) string {
	return "/" + path
}