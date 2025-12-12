package setup

import (
	notifController "github.com/NIROOZbx/project-server/internal/notification/controllers"
	notifRepo "github.com/NIROOZbx/project-server/internal/notification/repositories"
	notifService "github.com/NIROOZbx/project-server/internal/notification/services"

	"github.com/NIROOZbx/project-server/internal/shared/database"
)


func SetupDI() (*notifController.NotificationHandler,*notifService.NotificationService){
	var db = database.DB()

	 notifRepository:=notifRepo.NewNotificationRepo(db)
	 notificationService:=notifService.NewNotificationService(notifRepository)
	 notificationHandler:=notifController.NewNotifController(notificationService)


	 return notificationHandler, notificationService


}