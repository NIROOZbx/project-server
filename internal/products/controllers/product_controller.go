package controllers

import (
	"net/http"

	productServices "github.com/NIROOZbx/project-server/internal/products/services"
	"github.com/NIROOZbx/project-server/internal/shared/dtos"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)


func ShowSingleProductHandler(c *gin.Context) {

	productId := c.Param("id")

	product, err := productServices.ShowSingleProductService(productId)
	if err != nil {
		errMsg := err.Error()

		if errMsg == "invalid product ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
			return
		}

		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product data due to internal error."})
		return
	}

	c.JSON(http.StatusOK, product)

}

func ListProductsHandler(c *gin.Context) {

	var input dtos.PaginationInput

	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters. Page and Limit must be numbers."})
		return
	}

	response, err := productServices.GetPaginationResults(&input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products due to internal error."})
		return
	}

	c.JSON(http.StatusOK, response)

}

func GetHomeProducts(c *gin.Context) {

    products, err := productServices.GetHomePageProducts()

    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to fetch home products"})
        return
    }

    c.JSON(200, products)
}
