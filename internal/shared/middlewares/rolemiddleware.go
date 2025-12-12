package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(requiredRole string) gin.HandlerFunc{

	return func(c *gin.Context) {

		userRole:=c.GetString("role")

	if userRole!=requiredRole{
		c.JSON(http.StatusForbidden,gin.H{"message":"Unauthorized access"})
		c.Abort()
		return 
	}


	c.Next()

	}


}