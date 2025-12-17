package controllers

import (
	"fmt"
	"net/http"

	"github.com/NIROOZbx/project-server/internal/auth/models"
	"github.com/NIROOZbx/project-server/internal/auth/services"
	"github.com/NIROOZbx/project-server/internal/shared/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ForgotPasswordHandler
// --------------------------------------
// STEP 1: User submits email to request a password reset OTP.
// This handler does NOT reveal whether the email exists (security).
// It always sends a generic success response.
//

func ForgotPasswordHandler(c *gin.Context) {

	var otpReq models.OTPRequest

	if err := c.ShouldBindJSON(&otpReq); err != nil {
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

	err := services.ForgotPasswordService(c.Request.Context(), otpReq.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process request. Please try again later"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "If an account with that email exists, a password reset code has been sent."})

	
}

// VerifyResetOTPHandler
// --------------------------------------
// STEP 2: User submits email + OTP.
// If OTP is correct → generate and return a RESET TOKEN.
// This token allows only password reset (NOT login).
// --------------------------------------
func VerifyResetOTPHandler(c *gin.Context) {

	var verifyOTP models.VerifyOTPRequest

	if err := c.ShouldBindJSON(&verifyOTP); err != nil {
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

	resetToken, err := services.VerifyResetPasswordService(c.Request.Context(), verifyOTP.Email, verifyOTP.OTP)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reset-token": resetToken})

}

// ResetPasswordHandler
// --------------------------------------
// STEP 3: User sends:
//   1. reset_token (in Authentication header)
//   2. new_password + confirm_password
//
// Middleware (ResetTokenMiddleware) extracts userID from reset token.
// This handler receives that userID and updates the password.
// -

func ResetPasswordHandler(c *gin.Context) {

	var changePassData models.ForgotPassword

	if err := c.ShouldBindJSON(&changePassData); err != nil {
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

	userID := c.MustGet("userID").(uint)

	err := services.ChangePasswordService(userID, changePassData.NewPassword)
	fmt.Println("Error", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password. Please try again later."})
		return
	}

	c.SetCookie(cookieName, "", -1, "/", cookieDomain, false, true)

	c.JSON(200, gin.H{"message": "Password updated successfully"})

}
