package controllers

import (
	"net/http"

	addressModel "github.com/NIROOZbx/project-server/internal/address/models"
	addressService "github.com/NIROOZbx/project-server/internal/address/services"
	"github.com/NIROOZbx/project-server/internal/shared/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AddAddressHandler(c *gin.Context) {

	var addressInput addressModel.AddressInput
	if err := c.ShouldBindJSON(&addressInput); err != nil {
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

	err := addressService.AddAddressService(userId, addressInput)

	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cannot add address. User not found."})
			return
		}

		if err.Error() == "maximum address limit reached" {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return

		}

	}
	c.JSON(200,gin.H{"message": "Address added successfully"})

}

func GetAddressesHandler(c *gin.Context) {

	userId := c.GetUint("userID")

	addresses, err := addressService.GetAddressService(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve addresses"})
		return
	}

	c.JSON(http.StatusOK, addresses)

}

func DeleteAddressHandler(c *gin.Context) {

	userId := c.GetUint("userID")

	addressId := c.Param("id")

	err := addressService.DeleteAddressService(userId, addressId)

	if err != nil {

		if err.Error() == "invalid address ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address ID"})
			return
		}

		if err.Error() == "address not found or access denied" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete address"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address deleted successfully"})

}

func SetDefaultAddressHandler(c *gin.Context){

	userId := c.GetUint("userID")

	addressId := c.Param("id")

	err := addressService.UpdateDefaultAddressService(userId, addressId)

	if err != nil {

		if err.Error() == "invalid address ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address ID"})
			return
		}

		if err.Error() == "address not found or access denied" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set default address"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address set to default successfully"})

}
