package controller

import (
	"github.com/gin-gonic/gin"
	"go-api/internal/dto"
	"go-api/internal/helper"
	"go-api/internal/model"
	"go-api/internal/repository"
	"go-api/internal/service"
)

type ProfileController struct {
	ProfileService *service.ProfileService
	FileRepo       *repository.FileRepository
}

func NewProfileController(profileService *service.ProfileService, fileRepo *repository.FileRepository) *ProfileController {
	return &ProfileController{ProfileService: profileService, FileRepo: fileRepo}
}

func (ctl *ProfileController) Update(c *gin.Context) {
	authUser, exists := c.Get("auth_user")
	if !exists {
		helper.Error(c, 401, "unauthorized")
		return
	}

	user := authUser.(*model.User)

	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ValidationError(c, helper.FormatValidationError(err))
		return
	}

	profile, err := ctl.ProfileService.Update(user.ID, req)
	if err != nil {
		helper.Error(c, 400, err.Error())
		return
	}

	helper.Success(c, "profile updated successfully", profile)
}

func (ctl *ProfileController) Get(c *gin.Context) {
	authUser, exists := c.Get("auth_user")
	if !exists {
		helper.Error(c, 401, "unauthorized")
		return
	}

	user := authUser.(*model.User)

	profile, err := ctl.ProfileService.GetByUserID(user.ID)
	if err != nil {
		helper.Error(c, 404, "profile not found")
		return
	}

	var avatarURL *string
	if profile.AvatarFileID != nil {
		file, err := ctl.FileRepo.FindByID(*profile.AvatarFileID)
		if err == nil {
			url := helper.ResolveFileURL(file.Path)
			avatarURL = &url
		}
	}

	res := dto.ProfileResponse{
		ID:           profile.ID,
		UserID:       profile.UserID,
		AvatarFileID: profile.AvatarFileID,
		AvatarURL:    avatarURL,
		FirstName:    profile.FirstName,
		LastName:     profile.LastName,
		Phone:        profile.Phone,
		Gender:       profile.Gender,
		BirthDate:    profile.BirthDate.Format("2006-01-02"),
		Bio:          profile.Bio,
		Country:      profile.Country,
		Province:     profile.Province,
		City:         profile.City,
		District:     profile.District,
		PostalCode:   profile.PostalCode,
		Address:      profile.Address,
	}

	helper.Success(c, "success", res)
}