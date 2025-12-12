package main

import (
	"log"

	"github.com/NIROOZbx/project-server/internal/global"
	"github.com/NIROOZbx/project-server/internal/notification/setup"
	routes "github.com/NIROOZbx/project-server/internal/router"
	"github.com/NIROOZbx/project-server/internal/shared/cache"
	"github.com/NIROOZbx/project-server/internal/shared/config"
	"github.com/NIROOZbx/project-server/internal/shared/database"
	"github.com/NIROOZbx/project-server/internal/shared/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	database.InitDb()

	cache.InitRedis()

	database.RunMigrations()

	router := gin.Default()

	router.Use(middlewares.CORSMiddleware(),middlewares.RateLimiter())

	notifHandler,notifService:=setup.SetupDI()

	routes.RegisterAPIRoutes(router,notifHandler)

	global.NotifService=notifService

	log.Println("Server starting on port 8080...")

	router.Run()
}