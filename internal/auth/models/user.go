package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name          string `gorm:"type:varchar(100);not null;uniqueIndex" json:"name"`
	Password      string `gorm:"type:varchar(255);not null;" json:"password -"`
	Role          string `gorm:"type:varchar(50);default:'user'" json:"role"`
	Email         string `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	ProfileImage  string `gorm:"type:varchar(255)"`
	IsBlocked     bool   `gorm:"default:false" json:"is_blocked"`
	BlockedReason string `gorm:"type:varchar(255)" json:"blocked_reason"`
	TokenVersion  int    `gorm:"default:1;not null"`
	IsVerified    bool   `gorm:"default:false;not null"`
}
