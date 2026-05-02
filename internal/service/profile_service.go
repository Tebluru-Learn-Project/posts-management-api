package service

import (
	"time"

	"go-api/internal/dto"
	"go-api/internal/model"
	"go-api/internal/repository"
)

type ProfileService struct {
	ProfileRepo *repository.ProfileRepository
}

func NewProfileService(profileRepo *repository.ProfileRepository) *ProfileService {
	return &ProfileService{ProfileRepo: profileRepo}
}

func (s *ProfileService) Update(userID uint64, req dto.UpdateProfileRequest) (*model.Profile, error) {
	birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		return nil, err
	}

	// Check if profile already exists, if not create one
	profile, err := s.ProfileRepo.FindByUserID(userID)

	// profile belum ada → create
	if err != nil {
		newProfile := &model.Profile{
			UserID:     userID,
			FirstName:  req.FirstName,
			LastName:   req.LastName,
			Phone:      req.Phone,
			Gender:     req.Gender,
			BirthDate:  birthDate,
			Bio:        req.Bio,
			Country:    req.Country,
			Province:   req.Province,
			City:       req.City,
			District:   req.District,
			PostalCode: req.PostalCode,
			Address:    req.Address,
		}
		if req.AvatarFileID != nil {
			newProfile.AvatarFileID = req.AvatarFileID
		}

		if err := s.ProfileRepo.Create(newProfile); err != nil {
			return nil, err
		}

		return newProfile, nil
	}

	// update profile yang sudah ada
	profile.FirstName = req.FirstName
	profile.LastName = req.LastName
	profile.Phone = req.Phone
	if req.AvatarFileID != nil {
		profile.AvatarFileID = req.AvatarFileID
	}
	profile.Gender = req.Gender
	profile.BirthDate = birthDate
	profile.Bio = req.Bio
	profile.Country = req.Country
	profile.Province = req.Province
	profile.City = req.City
	profile.District = req.District
	profile.PostalCode = req.PostalCode
	profile.Address = req.Address

	if err := s.ProfileRepo.Update(profile); err != nil {
		return nil, err
	}

	return profile, nil
}

func (s *ProfileService) GetByUserID(userID uint64) (*model.Profile, error) {
	return s.ProfileRepo.FindByUserID(userID)
}
