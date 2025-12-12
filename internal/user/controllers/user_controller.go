package controllers

import (
	"fmt"
	"net/http"

	"github.com/NIROOZbx/project-server/internal/shared/dtos"
	"github.com/NIROOZbx/project-server/internal/shared/validation"
	userModel "github.com/NIROOZbx/project-server/internal/user/models"
	userService "github.com/NIROOZbx/project-server/internal/user/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ShowUserProfile(c *gin.Context) {

	userId := c.GetUint("userID")

	userProfile, err := userService.GetUserProfile(userId)

	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "User profile not found."})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve profile data due to internal error."})
		return
	}

	c.JSON(http.StatusOK, userProfile)

}

func UpdateUserName(c *gin.Context) {

	var newUserName userModel.ChangeName

	if err := c.ShouldBindJSON(&newUserName); err != nil {
		if validationError, ok := err.(validator.ValidationErrors); ok {
			errorMessages := make(map[string]string)

			for _, fe := range validationError {
				errorMessages[fe.Field()] = validation.FieldValidator(fe)
			}

			c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	userId := c.GetUint("userID")

	err := userService.UpdateUserName(userId, newUserName.Name)

	if err != nil {
		if err.Error() == "no user found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User profile not found."})
			return
		}
		if err.Error() == "exists" {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists,try another"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update username due to internal server error."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Username updated successfully."})

}

func UpdateUserPassword(c *gin.Context) {
	var newPassword userModel.ResetPassword

	if err := c.ShouldBindJSON(&newPassword); err != nil {
		if validationError, ok := err.(validator.ValidationErrors); ok {
			errorMessages := make(map[string]string)

			for _, fe := range validationError {
				errorMessages[fe.Field()] = validation.FieldValidator(fe)
			}

			c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	userId := c.GetUint("userID")

	err := userService.UpdatePassword(userId, newPassword.NewPassword, newPassword.OldPassword)

	if err != nil {
		if err.Error() == "the old password you entered is incorrect" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "cannot be same as old password" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password due to server error."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})

}

func GetAlluserHandler(c *gin.Context) {

	var paginationInput dtos.PaginationInput

	if err := c.ShouldBindQuery(&paginationInput); err != nil {
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

	allUser, err := userService.GetAllUserService(paginationInput.Page, paginationInput.Limit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, allUser)

}

func BlockUserHandler(c *gin.Context) {

	var input userModel.BlockUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input. Send {'is_blocked': true/false}"})
		return
	}

	userID := c.Param("id")

	adminId := c.GetUint("userID")

	err := userService.BlockUserService(userID, adminId, input.IsBlocked)

	if err != nil {
		errMsg := err.Error()

		fmt.Println("err", err)

		if errMsg == "invalid user ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
			return
		}

		if errMsg == "cannot block your own account" {
			c.JSON(http.StatusConflict, gin.H{"error": errMsg})
			return
		}

		if errMsg == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": errMsg})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to block user"})
		return
	}
	statusMsg := "unblocked"
	if *input.IsBlocked {
		statusMsg = "blocked"
	}

	c.JSON(http.StatusOK, gin.H{"message": "User " + statusMsg + " successfully"})

}

func UpdateUserProfileImage(c *gin.Context) {

	var profileInput userModel.UserProfileImage

	if err := c.ShouldBind(&profileInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image data"})
		return
	}

	userID := c.GetUint("userID")

	imageURL, err := userService.UploadProfileImage(c.Request.Context(), userID, profileInput.ProfileImage)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Profile image updated successfully",
		"image_url": imageURL,
	})

}
