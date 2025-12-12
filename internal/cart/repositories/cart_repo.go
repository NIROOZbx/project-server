package repositories

import (
	cartModel "github.com/NIROOZbx/project-server/internal/cart/models"
	productModel "github.com/NIROOZbx/project-server/internal/products/models"
	"github.com/NIROOZbx/project-server/internal/shared/database"
	"gorm.io/gorm"
)

func AddToCartInDB(userId uint, productId uint) error {

	var db = database.DB()

	newItem := cartModel.Cart{
		UserID:    userId,
		ProductID: productId,
	}

	res := db.Create(&newItem)

	return res.Error

}

func UpdateCartQuantityInDB(userId uint, productId uint, quantity int) error {

	var db = database.DB()

	result := db.Model(&cartModel.Cart{}).Where("user_id = ? AND product_id = ?", userId, productId).Update("quantity", quantity)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}

func RemoveFromCartInDB(userId uint, productId uint) error {

	var db = database.DB()

	res := db.Where("user_id = ? AND product_id = ?", userId, productId).Delete(&cartModel.Cart{})

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}

func GetCartFromDB(userId uint) ([]cartModel.Cart,error){

		var db = database.DB()

		var userCart []cartModel.Cart

		res:=db.Preload("Product").Where("user_id = ?",userId).Find(&userCart)


		return userCart,res.Error


}

func FindProductById(productId uint)(*productModel.Product,error){

	var db=database.DB()

	var product productModel.Product

	res:=db.First(&product,"id = ?",productId)


	return &product,res.Error



}