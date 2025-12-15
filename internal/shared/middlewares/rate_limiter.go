package middlewares

import (
	"context"
	"net/http"

	"github.com/NIROOZbx/project-server/internal/shared/cache"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
)

func RateLimiter() gin.HandlerFunc{

	return func (c *gin.Context) {
		clientIP := c.ClientIP()
		ctx := context.Background()

		limiter:=cache.GetRateLimiter()

		res,err:=limiter.Allow(ctx,clientIP,redis_rate.PerMinute(10))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Rate limit check failed"})
			c.Abort()
			return
		}
		if res.Allowed == 0 {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			c.Abort()
			return
		}
		c.Next()
	}

}