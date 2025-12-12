package models

import (
	"time"

	addressModel "github.com/NIROOZbx/project-server/internal/address/models"
	authModel "github.com/NIROOZbx/project-server/internal/auth/models"
	prodModel "github.com/NIROOZbx/project-server/internal/products/models"
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	OrderID         uint              `gorm:"not null;index" json:"order_id"`
	Order           Order             `gorm:"foreignKey:OrderID" references:"ID" json:"-"`
	ProductID       uint              `gorm:"not null" json:"product_id"`
	Product         prodModel.Product `gorm:"foreignKey:ProductID" json:"-"`
	PriceAtPurchase float64           `gorm:"not null" json:"price"`
	Quantity        int               `gorm:"not null" json:"quantity"`
	ProductName     string            `gorm:"type:varchar(255);not null" json:"product_name"`
	ImageURL        string            `gorm:"type:text" json:"image_url"`
	ItemStatus      string            `gorm:"type:varchar(20);default:'active'"`
	CancelledReason string            `gorm:"type:varchar(255);"`
}

type Order struct {
	gorm.Model

	UserID uint           `gorm:"not null;index" json:"user_id"`
	User   authModel.User `gorm:"foreignKey:UserID" json:"-"`

	AddressID uint                 `gorm:"not null" json:"address_id"`
	Address   addressModel.Address `gorm:"foreignKey:AddressID" json:"address"`

	TotalPrice    float64     `gorm:"not null" json:"total_price"`
	Status        string      `gorm:"type:varchar(20);default:'pending'" json:"status"`
	PaymentMethod string      `gorm:"type:varchar(20);not null" json:"payment_method"`
	OrderDate     time.Time   `gorm:"autoCreateTime;not null" json:"order_date"`
	Reference     string      `gorm:"type:varchar(50);uniqueIndex;not null" json:"order_reference"`
	Items         []OrderItem `gorm:"foreignKey:OrderID;references:ID;constraint:OnDelete:CASCADE" json:"items"`
}

type PlaceOrderInput struct {
	AddressID     uint   `json:"address_id" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required,oneof=COD RAZORPAY"`
}

type OrderItemDTO struct {
	ItemID          uint    `json:"item_id"`
	ProductID       uint    `json:"product_id"`
	ProductName     string  `json:"product_name"`
	ProductImageURL string  `json:"image_url"`
	Quantity        int     `json:"quantity"`
	PriceAtPurchase float64 `json:"price"`
	ItemStatus      string  `json:"item_status"`
}

type AddressDTO struct {
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	State         string `json:"state"`
	ZipCode       string `json:"zip_code"`
	Country       string `json:"country"`
}

type OrderResponse struct {
	OrderID uint `json:"order_id"`

	Status        string         `json:"status"`
	TotalPrice    float64        `json:"total_price"`
	OrderDate     time.Time      `json:"order_date"`
	Items         []OrderItemDTO `json:"items"`
	Address       AddressDTO     `json:"address"`
	Reference     string         `json:"Reference"`
	PaymentMethod string         `json:"payment_method"`
}

type OrderStatus struct {
	Status string `json:"status"`
}

type AdminOrderDTO struct {
	OrderID       uint      `json:"order_id"`
	ItemStatus    string    `json:"item_status"`
	ItemID        uint      `json:"item_id"`
	OrderDate     time.Time `json:"order_date"`
	CustomerName  string    `json:"customer_name"`
	ProductID     uint      `json:"product_id"`
	ProductName   string    `json:"product_name"`
	ProductImage  string    `json:"product_image"`
	ProductPrice  float64   `json:"product_price"`
	Quantity      int       `json:"quantity"`
	TotalPrice    float64   `json:"total_price"`
	CustomerEmail string    `json:"customer_email"`
	ProductLeague string    `json:"category"`
}

type OrderStats struct {
	TotalItems      int64
	Active          int64
	Shipped         int64
	Delivered       int64
	Cancelled       int64
	TotalRevenue    float64
	TotalProducts   int64
	TotalUsers      int64
	EstimatedProfit float64
}

type Result struct {
	ItemStatus string  `gorm:"column:item_status"`
	Count      int64   `gorm:"column:count"`
	TotalPrice float64 `gorm:"column:total_amount"`
}


type OrderCancelReasons struct{
	CancelledReason string `json:"reason"`
}


type AdminCancelReason struct{
	CancelledReason string `json:"reason"`
}