package models

import (
	"gorm.io/gorm"
	"mime/multipart"

)

type Product struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(255);not null" json:"name"`
	Team        string  `gorm:"type:varchar(255);not null" json:"team"`
	League      string  `gorm:"type:varchar(255);not null" json:"league"`
	Season      int     `gorm:"not null" json:"season"`
	Stock       int     `json:"stock"`
	Price       float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	Currency    string  `gorm:"type:char(3);not null;default:'$'" json:"currency"`
	Image       string  `gorm:"type:text;not null" json:"image"`
	Category    string  `gorm:"type:varchar(50);not null" json:"category"`
	Description string  `gorm:"type:text;not null" json:"description"`
}



type AddProductInput struct {
	Name        string                `form:"name" binding:"required,min=2"`
	Team        string                `form:"team" binding:"required"`
	League      string                `form:"league" binding:"required"`
	Season      int                   `form:"season" binding:"required"`
	Stock       int                   `form:"stock" binding:"required,min=0"`
	Price       float64               `form:"price" binding:"required,gt=0"`
	Currency    string                `form:"currency" binding:"required"`
	Category    string                `form:"category" binding:"required"`
	Description string                `form:"description" binding:"required"`

	ImageFile   *multipart.FileHeader `form:"image" binding:"required"` 
}

type UpdateProductInput struct {
	Name        string                `form:"name"` // No 'required'
	Team        string                `form:"team"`
	League      string                `form:"league"`
	Season      int                   `form:"season"`
	Stock       *int                  `form:"stock" binding:"omitempty,min=0"` 
	Price       float64               `form:"price" binding:"omitempty,gt=0"`
	Currency    string                `form:"currency"`
	Category    string                `form:"category"`
	Description string                `form:"description"`
	ImageFile   *multipart.FileHeader `form:"image"` 
}