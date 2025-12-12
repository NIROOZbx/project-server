package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model

	OrderID uint `gorm:"not null;index"`
	UserID  uint `gorm:"not null"`
	Amount        float64 `gorm:"not null"`
	PaymentMethod string  `gorm:"type:varchar(50);not null"`
	PaymentStatus string  `gorm:"type:varchar(20);default:'pending'"`
	TransactionID string  `gorm:"type:varchar(255)"` 
}