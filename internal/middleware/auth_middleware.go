package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go-api/internal/helper"
	"go-api/internal/repository"
)

func AuthMiddleware(
	sessionRepo *repository.SessionRepository,
	authRepo *repository.AuthRepository,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			helper.Error(c, 401, "unauthorized")
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			helper.Error(c, 401, "unauthorized")
			c.Abort()
			return
		}

		token := tokenParts[1]

		session, err := sessionRepo.FindByToken(token)
		if err != nil {
			helper.Error(c, 401, "unauthorized")
			c.Abort()
			return
		}

		if time.Now().After(session.ExpiredAt) {
			helper.Error(c, 401, "session expired")
			c.Abort()
			return
		}

		user, err := authRepo.FindUserByID(session.UserID)
		if err != nil {
			helper.Error(c, 401, "unauthorized")
			c.Abort()
			return
		}

		c.Set("auth_user", user)
		c.Next()
	}
}