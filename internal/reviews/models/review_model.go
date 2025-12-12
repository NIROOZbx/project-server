package models

import (
	"time"

	authModel "github.com/NIROOZbx/project-server/internal/auth/models"
	productModel "github.com/NIROOZbx/project-server/internal/products/models"
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID    uint                 `gorm:"not null;index" json:"user_id"`
	User      authModel.User       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
	ProductID uint                 `gorm:"not null;index" json:"product_id"`
	Product   productModel.Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Comment   string               `gorm:"type:text" json:"comment"`
}

type ReviewInput struct{
	Comment string `json:"comment"`
}


type ReviewResponse struct {
    ID        uint      `json:"id"`
    Comment   string    `json:"comment"`
    CreatedAt time.Time `json:"created_at"`
    UserName  string    `json:"user_name"` // We only send the name!
	UserImage string `json:"user_image"`
}