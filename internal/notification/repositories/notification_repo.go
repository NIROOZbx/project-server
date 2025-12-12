package repositories

import (
	notifModel "github.com/NIROOZbx/project-server/internal/notification/models"
	"gorm.io/gorm"
)

type NotificationRepo struct {
	DB *gorm.DB
}

func NewNotificationRepo(db *gorm.DB) *NotificationRepo {
	return &NotificationRepo{DB: db}
}

func (n *NotificationRepo) CreateBatch(notif []notifModel.Notification) error {

return n.DB.CreateInBatches(notif,100).Error
}


func (n *NotificationRepo) GetNotificationsFromDB(userID uint)([]notifModel.Notification,error){

	var notifications []notifModel.Notification

	res:=n.DB.Model(&notifModel.Notification{}).Where("user_id = ?",userID).Order("created_at DESC").Find(&notifications)

	return notifications,res.Error

}


func (n *NotificationRepo) MarkAllReadInDB(userID uint)error{

	err := n.DB.Model(&notifModel.Notification{}).
        Where("user_id = ? AND is_read = ?", userID, false).
        Update("is_read", true).Error
        
    return err

}

func (n *NotificationRepo) MarkReadInDB(notifID, userID uint) error {

    result := n.DB.Model(&notifModel.Notification{}).
        Where("id = ? AND user_id = ?", notifID, userID).
        Update("is_read", true)

    if result.Error != nil {
        return result.Error
    }

    if result.RowsAffected == 0 {
        return gorm.ErrRecordNotFound
    }

    return nil

}