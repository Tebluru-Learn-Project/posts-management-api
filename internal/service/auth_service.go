package service

import (
	"errors"
	"go-api/internal/dto"
	"go-api/internal/helper"
	"go-api/internal/model"
	"go-api/internal/repository"
	"time"
	"fmt"
)

type AuthService struct {
	AuthRepo *repository.AuthRepository
	OTPRepo  *repository.OTPRepository
	MailSvc  *MailService
}

func NewAuthService(
	authRepo *repository.AuthRepository,
	otpRepo *repository.OTPRepository,
	mailSvc *MailService,
	) *AuthService {
	return &AuthService{
		AuthRepo: authRepo,
		OTPRepo:  otpRepo,
		MailSvc:  mailSvc,
	}
}

func (s *AuthService) Login(req *dto.LoginRequest, ipAddress, userAgent string) (*dto.LoginResponse, error) {
	user, err := s.AuthRepo.FindUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("Invalid email")
	}
	
	if !user.IsActive {
		return nil, errors.New("account is not active")
	}

	if !helper.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("Invalid password")
	}
	
	token, err := helper.GenerateToken(32)
	if err != nil {
		return nil, err
	}
	
	session := &model.Session{
		UserID: user.ID,
		Token: token,
		IPAddress: &ipAddress,
		UserAgent: &userAgent,
		ExpiredAt: time.Now().Add(24 * time.Hour),
		LastActivity: time.Now(),
	}
	
	if err := s.AuthRepo.CreateSession(session); err != nil {
		// log error
		fmt.Println("error creating session:", err)
		return nil, errors.New("failed to create session")
	}
	
	return &dto.LoginResponse{Token: token}, nil
}

func (s *AuthService) Register(req dto.RegisterRequest) error {
	existingUser, err := s.AuthRepo.FindUserByEmail(req.Email)
	if err == nil && existingUser != nil {
		return errors.New("email already registered")
	}

	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		IsActive: false,
	}

	if err := s.AuthRepo.CreateUser(user); err != nil {
		return err
	}

	code := helper.GenerateOTP()

	otp := &model.OTP{
		UserID:    user.ID,
		Code:      code,
		Purpose:   "register",
		ExpiredAt: time.Now().Add(5 * time.Minute),
	}

	if err := s.OTPRepo.CreateOTP(otp); err != nil {
		return err
	}

	if err := s.MailSvc.SendOTP(user.Email, code); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) VerifyOTP(req dto.VerifyOTPRequest) error {
	user, err := s.AuthRepo.FindUserByEmail(req.Email)
	if err != nil {
		return errors.New("Invalid email request")
	}

	otp, err := s.OTPRepo.FindByUserIDAndCode(user.ID, req.OTP)
	if err != nil {
		return errors.New("invalid or expired otp")
	}

	if time.Now().After(otp.ExpiredAt) {
		_ = s.OTPRepo.DeleteByID(otp.ID)
		return errors.New("otp has expired")
	}

	if err := s.AuthRepo.ActivateUser(user.ID); err != nil {
		return err
	}

	if err := s.OTPRepo.DeleteByID(otp.ID); err != nil {
		return err
	}

	return nil
}