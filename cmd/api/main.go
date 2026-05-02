package main

import (
	"log"
	"go-api/config"
	"go-api/internal/controller"
	"go-api/internal/repository"
	"go-api/internal/route"
	"go-api/internal/service"
	"go-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment
	config.LoadEnv()

	// Load app config
	appConfig := config.LoadAppConfig()

	// Connect database
	config.ConnectDB()

	// Init gin
	r := gin.Default()

	// Init repositories
	authRepo := repository.NewAuthRepository(config.DB)
	otpRepo := repository.NewOTPRepository(config.DB)

	// Init services
	mailService := service.NewMailService()
	authService := service.NewAuthService(authRepo, otpRepo, mailService)

	// Init controller
	authController := controller.NewAuthController(authService)
	// Session
	sessionRepo := repository.NewSessionRepository(config.DB)
	// Init middleware
	authMiddleware := middleware.AuthMiddleware(sessionRepo, authRepo)

	// Init upload
	uploadService := service.NewUploadService()
	fileRepo := repository.NewFileRepository(config.DB)
	uploadController := controller.NewUploadController(uploadService, fileRepo)

	// Init profile
	profileRepo := repository.NewProfileRepository(config.DB)
	profileService := service.NewProfileService(profileRepo)
	profileController := controller.NewProfileController(profileService, fileRepo)


	// Setup routes
	route.SetupAPI(r, authController, profileController, uploadController, authMiddleware)

	// Run server
	log.Printf("%s running on port %s", appConfig.AppName, appConfig.AppPort)
	if err := r.Run(":" + appConfig.AppPort); err != nil {
		log.Fatal("failed to start server:", err)
	}
}