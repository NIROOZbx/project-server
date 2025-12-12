package middlewares

import (
	"fmt"

	"net/http"
	"strings"

	"github.com/NIROOZbx/project-server/internal/auth/repositories"
	"github.com/NIROOZbx/project-server/internal/shared/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		secret := config.GetConfig().AccessSecret

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		fmt.Println("Access token",tokenString)

		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token format"})
			c.Abort()
			return
		}

		accessToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return []byte(secret), nil
		})

		if err != nil || !accessToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized: Invalid token"})
			c.Abort()
			return
		}

		fullToken, ok := accessToken.Claims.(jwt.MapClaims)

		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims structure"})
			return
		}

		userID := uint(fullToken["userID"].(float64))
		role := fullToken["Role"].(string)
		version := int(fullToken["version"].(float64))

		user, DbErr := repositories.FindUserById(userID)

		if DbErr != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}

		if user.TokenVersion != version {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Session expired, please log in again"})
			return
		}

		if user.IsBlocked{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Blocked by admin"})
			return

		}

		fmt.Println("Role", role)

		c.Set("userID", userID)
		c.Set("role", role)

		c.Next()

	}
}
