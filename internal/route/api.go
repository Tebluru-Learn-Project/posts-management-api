package route

import (
	"go-api/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetupAPI(r *gin.Engine, authController *controller.AuthController, 
	profileController *controller.ProfileController, 
	uploadController *controller.UploadController,
	authMiddleware gin.HandlerFunc) {
	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authController.Login)
			auth.POST("/register", authController.Register)
			auth.POST("/verify-otp", authController.VerifyOTP)
		}
		user := api.Group("/user")
		user.Use(authMiddleware)
		{
			user.GET("/me", authController.Me)
			user.GET("/profile", profileController.Get)
			user.PUT("/profile", profileController.Update)
			user.POST("/avatar", uploadController.UploadAvatar)
		}
	}
	
}
