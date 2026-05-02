package controller

import (
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"go-api/internal/helper"
	"go-api/internal/model"
	"go-api/internal/repository"
	"go-api/internal/service"
)

type UploadController struct {
	UploadService *service.UploadService
	FileRepo      *repository.FileRepository
}

func NewUploadController(
	uploadService *service.UploadService,
	fileRepo *repository.FileRepository,
) *UploadController {
	return &UploadController{
		UploadService: uploadService,
		FileRepo:      fileRepo,
	}
}

func (ctl *UploadController) UploadAvatar(c *gin.Context) {
	authUser, exists := c.Get("auth_user")
	if !exists {
		helper.Error(c, 401, "unauthorized")
		return
	}

	user := authUser.(*model.User)

	file, err := c.FormFile("avatar")
	if err != nil {
		helper.Error(c, 400, "avatar file is required")
		return
	}

	fileName, path, err := ctl.UploadService.SaveImage(file, "avatars")
	if err != nil {
		helper.Error(c, 400, err.Error())
		return
	}

	ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(file.Filename)), ".")

	uploadedFile := &model.File{
		UserID:       user.ID,
		Disk:         "local",
		Path:         path,
		Filename:     fileName,
		OriginalName: file.Filename,
		MimeType:     file.Header.Get("Content-Type"),
		Extension:    ext,
		Size:         uint64(file.Size),
		Visibility:   "public",
		Category:     "avatar",
		IsUsed:       false,
	}

	if err := ctl.FileRepo.Create(uploadedFile); err != nil {
		helper.Error(c, 500, "failed to save file metadata")
		return
	}

	helper.Success(c, "avatar uploaded successfully", gin.H{
		"file_id": uploadedFile.ID,
	})
}