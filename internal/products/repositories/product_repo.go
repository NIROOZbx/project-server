package repositories

import (
	productModel "github.com/NIROOZbx/project-server/internal/products/models"
	"github.com/NIROOZbx/project-server/internal/shared/database"
	"github.com/NIROOZbx/project-server/internal/shared/dtos"
	"gorm.io/gorm"
)

func GetSingleProductFromDB(productId uint) (*productModel.Product, error) {

	var db = database.DB()

	var product productModel.Product

	res := db.First(&product, productId)

	if res.Error != nil {
		return nil, res.Error
	}

	return &product, nil

}

func GetPaginationFromDB(offset int, input *dtos.PaginationInput) ([]productModel.Product, error) {

	db := database.DB()
	var products []productModel.Product

	query := db.Model(&productModel.Product{})

	if input.Filter != "" {
		query.Where("league ILIKE ?", "%"+input.Filter+"%")
	}

	if input.Se != "" {
		query = query.Where("name ILIKE ?", "%"+input.Se+"%")
	}

	switch input.SortBy {

	case "popular":
		query = query.
			Select("products.*, COUNT(order_items.id) as sales_count").             
			Joins("LEFT JOIN order_items ON order_items.product_id = products.id"). 
			Group("products.id").                                                  
			Order("sales_count DESC")

	case "low":
		query = query.Order("price ASC")
	case "high":
		query = query.Order("price DESC")
	case "low-stock":
		query = query.Order("stock ASC")
	case "high-stock":
		query = query.Order("stock DESC")
	case "recent":
		query = query.Order("created_at ASC")
	case "asc":
		query = query.Order("name ASC")
	case "desc":
		query = query.Order("name DESC")
	default:
		query = query.Order("created_at")
	}

	query = query.Offset(offset).Limit(input.Limit)

	err := query.Find(&products).Error

	return products, err

}

func CountAllProducts() (int64, error) {
	db := database.DB()
	var count int64

	if err := db.Model(&productModel.Product{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func AddProductInDB(newProduct *productModel.Product) error {

	var db = database.DB()

	return db.Create(newProduct).Error

}

func RemoveProductFromDB(productId uint) error {

	var db = database.DB()

	res := db.Where("id = ?", productId).Delete(&productModel.Product{})

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}

func GetAllProductsWithPagination(limit int, offset int) ([]productModel.Product, error) {
	db := database.DB()
	var products []productModel.Product

	result := db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&products)

	return products, result.Error
}

func UpdateProductInDB(updatedProducts *productModel.Product) error {
	db := database.DB()
	result := db.Save(updatedProducts)

	if result.Error != nil {
		return result.Error
	}

	return nil

}
func GetLimitedProducts(limit int) ([]productModel.Product, error) {
	var products []productModel.Product
	result := database.DB().Limit(limit).Where("season = ?", 2025).Find(&products)

	return products, result.Error
}
