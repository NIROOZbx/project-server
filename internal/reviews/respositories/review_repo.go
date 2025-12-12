package repositories

import (
	reviewModel "github.com/NIROOZbx/project-server/internal/reviews/models"
	"github.com/NIROOZbx/project-server/internal/shared/database"
)

func AddProductReviewInDB(newReview *reviewModel.Review) error {

	var db = database.DB()

	return db.Create(newReview).Error

}

func HasUserPurchasedProduct(userID uint, productID uint) bool {
    db := database.DB()
    var count int64

    db.Table("order_items").
        Joins("JOIN orders ON orders.id = order_items.order_id").
        Where("orders.user_id = ? AND order_items.product_id = ? AND order_items.item_status = ? OR orders.status = ?", userID, productID, "delivered", "delivered").
        Count(&count)

    return count > 0
}

func GetReviewsByProductID(productID uint) ([]reviewModel.Review, error) {
	db := database.DB()
	var reviews []reviewModel.Review


	err := db.Preload("User").
		Where("product_id = ?", productID).
		Order("created_at desc").
		Find(&reviews).Error

	return reviews, err
}