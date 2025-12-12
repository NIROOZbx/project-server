package controllers

import (
	"net/http"

	"github.com/NIROOZbx/project-server/internal/auth/models"
	"github.com/NIROOZbx/project-server/internal/auth/services"
	"github.com/NIROOZbx/project-server/internal/shared/validation"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)


func VerifyOTPHandler(c *gin.Context){

	var VerifyOTP models.VerifyOTPRequest

	if err:=c.ShouldBindJSON(&VerifyOTP); err!=nil{
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

	err:=services.VerifyOTP(c.Request.Context(),VerifyOTP.Email,VerifyOTP.OTP)

	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired OTP",})
		return
	}

	c.JSON(http.StatusAccepted,gin.H{"message":"Email verified successfully"})

}

func ResendOTPHandler(c *gin.Context) {
	var OTPRequest models.OTPRequest

	if err:=c.ShouldBindJSON(&OTPRequest); err!=nil{
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
	

	err:=services.SendOTP(c.Request.Context(),OTPRequest.Email)	

	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP. Please try again later."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "If an account with that email exists, a new code has been sent."})
}