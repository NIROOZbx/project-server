package database

import (
	"log"

	addressModel "github.com/NIROOZbx/project-server/internal/address/models"
	authModel "github.com/NIROOZbx/project-server/internal/auth/models"
	cartModel "github.com/NIROOZbx/project-server/internal/cart/models"
	notifModel "github.com/NIROOZbx/project-server/internal/notification/models"
	orderModel "github.com/NIROOZbx/project-server/internal/order/models"
	productmodel "github.com/NIROOZbx/project-server/internal/products/models"
	reviewModel "github.com/NIROOZbx/project-server/internal/reviews/models"
	wishlistModel "github.com/NIROOZbx/project-server/internal/wishlist/models"
)

func RunMigrations() {
	db := DB()

	err := db.AutoMigrate(&authModel.User{},
		&productmodel.Product{},
		&cartModel.Cart{},
		&wishlistModel.Wishlist{},
		&addressModel.Address{},
		&orderModel.Order{},
		&orderModel.OrderItem{},
		&orderModel.Payment{},
		&reviewModel.Review{},
		&notifModel.Notification{},
	)

	if err!=nil{
		log.Fatal("AutoMigration failed:", err)
	}
}