package controller

import (
	"net/http"

	cartModel "github.com/NIROOZbx/project-server/internal/cart/models"
	cartServices "github.com/NIROOZbx/project-server/internal/cart/services"
	"github.com/NIROOZbx/project-server/internal/shared/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func AddToCartHandler(c *gin.Context) {
	var cartItems cartModel.AddToCartInput

	if err := c.ShouldBindJSON(&cartItems); err != nil {
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			errorMessages := make(map[string]string)

			for _, fe := range validationErr {
				errorMessages[fe.Field()] = validation.FieldValidator(fe)
			}

			c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
			return

		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	userId := c.GetUint("userID")

	err := cartServices.AddToCartService(userId, cartItems.ProductID)

	if err != nil {
		errMsg := err.Error()

		if errMsg == "product already exists in the cart" {

			c.JSON(http.StatusConflict, gin.H{"error": errMsg})
			return
		}

		if errMsg == "product does not exist" {
			c.JSON(http.StatusNotFound, gin.H{"error": errMsg})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item to cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added to cart successfully"})

}

func GetCartHandler(c *gin.Context) {

	userId := c.GetUint("userID")

	cartResponse, err := cartServices.GetCartService(userId)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cart details"})
		return
	}

	c.JSON(http.StatusOK, cartResponse)

}

func UpdateCartQuantityHandler(c *gin.Context) {
	var updateCart cartModel.UpdateCart

	if err := c.ShouldBindJSON(&updateCart); err != nil {
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			errorMessages := make(map[string]string)

			for _, fe := range validationErr {
				errorMessages[fe.Field()] = validation.FieldValidator(fe)
			}

			c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
			return

		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	userId := c.GetUint("userID")

	err := cartServices.UpdateCartQuantityService(userId, updateCart.ProductID, updateCart.Quantity)

	if err != nil {

		errMsg:=err.Error()

		if errMsg== "item not found in cart" || errMsg == "product not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found in cart"})
			return
		}
		if errMsg=="maximum quantity reached"{
			c.JSON(http.StatusBadRequest, gin.H{"error":errMsg})
			return
		}
		if errMsg=="no stock left"{
			c.JSON(http.StatusConflict, gin.H{"error":errMsg})
			return

		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart due to internal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart updated successfully"})

}

func RemoveFromCartHandler(c *gin.Context) {

	productId := c.Param("id")

	userId := c.GetUint("userID")

	err := cartServices.RemoveFromCartService(userId, productId)

	if err != nil {

		if err.Error() == "invalid product ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found in cart"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item due to internal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart successfully"})
}
