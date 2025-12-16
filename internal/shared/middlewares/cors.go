package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {

	config:=cors.DefaultConfig()

	config.AllowOrigins = []string{"http://localhost:5173","https://e-commerce-project-six-kappa.vercel.app"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept"}
	config.AllowCredentials = true

	return cors.New(config)

}