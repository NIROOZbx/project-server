package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RateLimiter() gin.HandlerFunc{


	return func (c *gin.Context) {
		clientIP:=c.ClientIP()

		fmt.Println("Client ip",clientIP)

		c.Next()
	}

}