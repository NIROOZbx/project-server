package controllers

import (
	"fmt"
	"net/http"

	productModel "github.com/NIROOZbx/project-server/internal/products/models"
	productServices "github.com/NIROOZbx/project-server/internal/products/services"
	"github.com/NIROOZbx/project-server/internal/shared/dtos"
	"github.com/NIROOZbx/project-server/internal/shared/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)



func AddProductHandler(c *gin.Context) {

	var prodInput productModel.AddProductInput
	if err := c.ShouldBind(&prodInput); err != nil {
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			errorMessages := make(map[string]string)

			for _, fe := range validationErr {

				errorMessages[fe.Field()] = validation.FieldValidator(fe)
			}

			c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters."})
		return
	}

	err := productServices.AddProductService(c.Request.Context(),&prodInput)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}


func DeleteProductHandler(c *gin.Context){

	productId:=c.Param("id")


	err:=productServices.DeleteProductService(productId)
	if err != nil {
		
		if err.Error() == "couldn't find the product" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

        if err.Error() == "invalid product ID format" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
            return
        }

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}


	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})


}


func GetAllProductsHandler(c * gin.Context){
	var input dtos.PaginationInput

	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	response, err := productServices.GetPaginationResults(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, response)

}

func UpdateProductHandler(c *gin.Context){
	productIDStr := c.Param("id")

	var input productModel.UpdateProductInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	err:=productServices.UpdateProductService(c.Request.Context(),productIDStr,&input)
	if err != nil {
		fmt.Println("Error in updating pordutc",err)
        if err == gorm.ErrRecordNotFound {
             c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
             return
        }
		if err.Error()=="invalid query format"{
			c.JSON(http.StatusNotFound, gin.H{"error":err.Error()})
             return
		}
		if err.Error()=="failed to upload new image"{
			c.JSON(http.StatusBadGateway, gin.H{"error":err.Error()})
            return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})

}

