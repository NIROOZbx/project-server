package models

import (
	userModel "github.com/NIROOZbx/project-server/internal/auth/models"
	productModel "github.com/NIROOZbx/project-server/internal/products/models"
)

type Cart struct {
	UserID    uint                 `gorm:"primaryKey;not null"`
	ProductID uint                 `gorm:"primaryKey;not null"`
	Quantity  int                  `gorm:"default:1;not null"`
	User      userModel.User       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Product   productModel.Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type UpdateCart struct {
	ProductID uint `json:"product_id" binding:"required,min=1"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

type AddToCartInput struct {
	ProductID uint `json:"product_id" binding:"required,min=1"`
}

type CartItemDTO struct {
	ID uint    `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	League    string  `json:"league"`
	ImageURL  string  `json:"image_url"`
	Quantity  int     `json:"quantity"`
	SubTotal  float64 `json:"sub_total"`
}

type CartResponseDTO struct {
	Items      []CartItemDTO `json:"items"`
	TotalItems int           `json:"total_items"`
	TotalPrice float64       `json:"total_price"`
}
