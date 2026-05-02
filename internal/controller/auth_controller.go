package controller

import (
	"go-api/internal/dto"
	"go-api/internal/helper"
	"go-api/internal/service"
	"go-api/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.ValidationError(ctx, helper.FormatValidationError(err))
		return
	}
	
	ip := ctx.ClientIP()
	userAgent := ctx.Request.UserAgent()
	
	result, err := c.AuthService.Login(&req, ip, userAgent)
	if err != nil {
		helper.Error(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	
	helper.Success(ctx, "Login successful", result)
}


func (ctl *AuthController) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ValidationError(c, helper.FormatValidationError(err))
		return
	}

	err := ctl.AuthService.Register(req)
	if err != nil {
		helper.Error(c, 400, err.Error())
		return
	}

	helper.Success(c, "verification code sent to your email", nil)
}

func (ctl *AuthController) VerifyOTP(c *gin.Context) {
	var req dto.VerifyOTPRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ValidationError(c, helper.FormatValidationError(err))
		return
	}

	err := ctl.AuthService.VerifyOTP(req)
	if err != nil {
		helper.Error(c, 400, err.Error())
		return
	}

	helper.Success(c, "account verified successfully", nil)
}

func (ctl *AuthController) Me(c *gin.Context) {
	authUser, exists := c.Get("auth_user")
	if !exists {
		helper.Error(c, 401, "unauthorized")
		return
	}

	user, ok := authUser.(*model.User)
	if !ok {
		helper.Error(c, 401, "unauthorized")
		return
	}

	role := ""
	if user.Role != nil {
		role = *user.Role
	}

	res := dto.MeResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Role:     role,
		IsActive: user.IsActive,
	}

	helper.Success(c, "success", res)
}
