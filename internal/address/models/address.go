package models

import (
	"gorm.io/gorm"
	userModels "github.com/NIROOZbx/project-server/internal/auth/models" 
)

type Address struct {
	gorm.Model
	UserID uint            `gorm:"not null;index"`
	User   userModels.User `gorm:"foreignKey:UserID;references:ID"`
	Name          string `gorm:"type:varchar(100);not null"`
	Phone         string `gorm:"type:varchar(20);not null"`
	StreetAddress string `gorm:"type:varchar(255);not null"`
	City          string `gorm:"type:varchar(100);not null"`
	State         string `gorm:"type:varchar(100);not null"`
	ZipCode       string `gorm:"type:varchar(20);not null"`
	Country       string `gorm:"type:varchar(100);not null"`
	IsDefault bool `gorm:"default:false"`
}


type AddressInput struct {
	Name          string `json:"name" binding:"required"`
	Phone         string `json:"phone" binding:"required"`
	StreetAddress string `json:"street_address" binding:"required"`
	City          string `json:"city" binding:"required"`
	State         string `json:"state" binding:"required"`
	ZipCode       string `json:"zip_code" binding:"required"`
	Country       string `json:"country" binding:"required"`
	IsDefault     bool   `json:"is_default"`
}

type AddressResponseDTO struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	State         string `json:"state"`
	ZipCode       string `json:"zip_code"`
	Country       string `json:"country"`
	IsDefault     bool   `json:"is_default"`
}

