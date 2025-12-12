package controllers

import (
	"net/http"

	"github.com/NIROOZbx/project-server/internal/shared/validation"
	wishlistModel "github.com/NIROOZbx/project-server/internal/wishlist/models"
	wishlistServices "github.com/NIROOZbx/project-server/internal/wishlist/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AddToWishlistHandler(c *gin.Context) {

	var wishlist wishlistModel.AddToWishlist

	if err := c.ShouldBindJSON(&wishlist); err != nil {
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

	err := wishlistServices.AddToWishlistService(userId, wishlist.ProductID)
	if err != nil {

		if err.Error() == "product already in wishlist" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "product not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to wishlist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added to wishlist"})
}

func RemoveWishlistHandler(c *gin.Context) {

	productId:=c.Param("id")

	userId := c.GetUint("userID")

	err := wishlistServices.RemoveFromWishlistService(userId,productId)

	if err != nil {
		if err.Error()=="invalid product ID format"{
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		if err.Error() == "product not found in wishlist" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove from wishlist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product removed from wishlist"})
}

func GetWishlistHandler(c *gin.Context) {

	userId := c.GetUint("userID")

	wishlist, err := wishlistServices.GetWishlistService(userId)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve wishlist"})
		return
	}
	c.JSON(http.StatusOK, wishlist)
}
