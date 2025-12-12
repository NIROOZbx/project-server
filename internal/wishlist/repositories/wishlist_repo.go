package repositories

import (
	"github.com/NIROOZbx/project-server/internal/shared/database"
	wishlistModel "github.com/NIROOZbx/project-server/internal/wishlist/models"
	"gorm.io/gorm"
)

func AddToWishlistInDB(userId uint, productId uint) error {

	var db = database.DB()

	newWishlist := &wishlistModel.Wishlist{
		UserID:    userId,
		ProductID: productId,
	}

	err := db.Create(newWishlist).Error

	if err != nil {
		return err
	}

	return nil

}

func RemoveFromWishlistInDB(userId uint, productId uint) error {

	var db = database.DB()

	res := db.Where("user_id = ? AND product_id = ?", userId, productId).Delete(&wishlistModel.Wishlist{})

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}

func GetWishlistFromDB(userId uint) ([]wishlistModel.Wishlist, error) {
	var db = database.DB()

	var wishlistItems []wishlistModel.Wishlist

	res := db.Preload("Product").Where("user_id=?", userId).Find(&wishlistItems)

	return wishlistItems, res.Error
}
