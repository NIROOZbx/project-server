package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
    gorm.Model
    UserID  uint   `gorm:"not null;index" json:"user_id"`
    Title   string `gorm:"type:varchar(100);not null" json:"title"`
    Message string `gorm:"type:text;not null" json:"message"`
	Date    time.Time `gorm:"autoCreateTime" json:"date"`
    IsRead  bool   `gorm:"default:false" json:"is_read"`
}

type CreateNotificationInput struct {
    UserID  []uint   `json:"user_id" binding:"required"` 
    Title   string `json:"title" binding:"required"`
    Message string `json:"message" binding:"required"`
 
}


type NotificationResponse struct {
    ID        uint      `json:"id"`
    Title     string    `json:"title"`
    Message   string    `json:"message"`
    Date      time.Time `json:"date"`    
    IsRead    bool      `json:"is_read"` 
}