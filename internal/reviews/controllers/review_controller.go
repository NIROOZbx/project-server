package controllers

import (
	"fmt"
	"net/http"

	reviewModel "github.com/NIROOZbx/project-server/internal/reviews/models"
	reviewService "github.com/NIROOZbx/project-server/internal/reviews/services"
	"github.com/gin-gonic/gin"
)

func AddReviewHandler(c *gin.Context) {

	var input reviewModel.ReviewInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	productID := c.Param("id")

	userID := c.GetUint("userID")

	err:=reviewService.AddProductReviewService(userID, productID, input.Comment)

	if err != nil {
		fmt.Println("Error in adding review",err)
		if err.Error()=="you can only review products you have purchased"{
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add review"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Review added successfully"})

}

func GetProductReviewsHandler(c *gin.Context) {
	productID := c.Param("id")

	reviews, err := reviewService.GetProductReviewsService(productID)

	if err != nil {
		if err.Error() == "invalid product ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reviews"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Reviews fetched successfully",
		"reviews": reviews,
	})
}