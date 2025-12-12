package middlewares

import (
	"net/http"
	"strings"

	"github.com/NIROOZbx/project-server/internal/shared/utils"
	"github.com/gin-gonic/gin"
)

func RequireResetToken() gin.HandlerFunc{
	return func(c *gin.Context) {
		authHeader:=c.GetHeader("Authorization")

		if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            return
        }

		tokenString:=strings.TrimPrefix(authHeader,"Bearer ")

		userID,err:=utils.ParseResetToken(tokenString)

		if err!=nil{
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired reset token"})
			return 
		}
		

		c.Set("userID",userID)
		c.Next()
	}
}
