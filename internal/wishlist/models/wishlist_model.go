package models

import (
	productModel "github.com/NIROOZbx/project-server/internal/products/models"
	userModel "github.com/NIROOZbx/project-server/internal/auth/models"

)

type Wishlist struct {
	UserID    uint `gorm:"primaryKey;not null"`
	ProductID uint `gorm:"primaryKey;not null"`
	User     userModel.User       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Product  productModel.Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type AddToWishlist struct {
	ProductID uint `json:"product_id" binding:"required,min=1"`
}

type WishlistItemDTO struct {
	Id   uint    `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
    League string `json:"league"`
}