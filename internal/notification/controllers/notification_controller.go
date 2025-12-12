package controllers

import (
	"net/http"

	notifModel "github.com/NIROOZbx/project-server/internal/notification/models"
	notifService "github.com/NIROOZbx/project-server/internal/notification/services"
	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	Service *notifService.NotificationService
}

func NewNotifController(service *notifService.NotificationService) *NotificationHandler {
	return &NotificationHandler{Service: service}
}

func (n *NotificationHandler) AddNotifcationHandler(c *gin.Context) {

	var input notifModel.CreateNotificationInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return

	}

	err := n.Service.SendBatch(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create notification"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Notification sent successfully"})
}


func (n *NotificationHandler) GetNotifications(c *gin.Context) {

	userId := c.GetUint("userID")

	notifications, err := n.Service.GetNotificationsService(userId)

	if err != nil {
        errMsg := err.Error()

        if errMsg == "invalid user ID format" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
            return
        }

        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications"})
        return
    }

	c.JSON(http.StatusOK, gin.H{
		"message":       "Success",
		"notifications": notifications,
	})

}


func (n *NotificationHandler) MarkAllRead(c *gin.Context){
	userID:=c.GetUint("userID")

	err:=n.Service.MarkAllReadService(userID)

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update notifications"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "All notifications marked as read"})
}

func (n *NotificationHandler) MarkRead(c *gin.Context) {

    userID := c.GetUint("userID")
    

    notifID := c.Param("id")

    err := n.Service.MarkReadService(notifID, userID)

    if err != nil {
        if err.Error() == "record not found" {
            c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
            return
        }
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update notification"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Notification marked as read"})
}