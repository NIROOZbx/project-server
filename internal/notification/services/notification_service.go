package services

import (
	"errors"
	"strconv"

	notifModel "github.com/NIROOZbx/project-server/internal/notification/models"
	notifRepo "github.com/NIROOZbx/project-server/internal/notification/repositories"
)

type NotificationService struct {
	Repo *notifRepo.NotificationRepo
}

func NewNotificationService(repo *notifRepo.NotificationRepo) *NotificationService {
	return &NotificationService{Repo: repo}
}

func (n *NotificationService) SendBatch(notifInput notifModel.CreateNotificationInput) error {

	notifications := make([]notifModel.Notification, 0, len(notifInput.UserID))
	for _, userID := range notifInput.UserID {
		notifications = append(notifications, notifModel.Notification{
			UserID:  userID,
			Title:   notifInput.Title,
			Message: notifInput.Message,
		})
	}

	return n.Repo.CreateBatch(notifications)

}


func (n *NotificationService) SendNotificationService(userID uint, title, message string) error {
    
 
    notif := notifModel.Notification{
        UserID:  userID,
        Title:   title,
        Message: message,
        IsRead:  false,
        
    }
    return n.Repo.CreateBatch([]notifModel.Notification{notif})
}

func (n *NotificationService) GetNotificationsService(userID uint) ([]notifModel.NotificationResponse, error) {

	dbNotifs, err := n.Repo.GetNotificationsFromDB(uint(userID))
	if err != nil {
		return nil, err
	}

	var response []notifModel.NotificationResponse

	for _, item := range dbNotifs {
		response = append(response, notifModel.NotificationResponse{
			ID:      item.ID,
			Title:   item.Title,
			Message: item.Message,
			Date:    item.CreatedAt,
			IsRead:  item.IsRead,
		})
	}

	if response == nil {
		response = []notifModel.NotificationResponse{}
	}

	return response, nil

}


func (n *NotificationService) MarkAllReadService(userID uint) error{


	return n.Repo.MarkAllReadInDB(userID)

}

func (n *NotificationService) MarkReadService(notifIdStr string, userID uint) error {
    notifID, err := strconv.Atoi(notifIdStr)
    if err != nil {
        return errors.New("invalid notification ID")
    }

    return n.Repo.MarkReadInDB(uint(notifID), userID)
}