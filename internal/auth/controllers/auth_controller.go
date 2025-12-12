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

// Constants: These define the shared values for cookies.
const (
	cookieName   = "user_session"
	cookieMaxAge = 60 * 60 * 24 * 7 // 7 days in seconds
	cookieDomain = "localhost"
)

// FieldValidator translates Gin/Validator error tags (e.g., "min", "required")
// into user-friendly strings for the frontend to display.

// RegisterHandler handles the POST /api/auth/signup request.
// RESPONSIBILITY: Validates input, calls the service, and handles conflict (409) or success (201).

func RegisterHandler(c *gin.Context) {

	var userinput models.SignInData

	if err := c.ShouldBindJSON(&userinput); err != nil {
		if validationError, ok := err.(validator.ValidationErrors); ok {
			errorMessages := make(map[string]string)

			for _, fe := range validationError {
				errorMessages[fe.Field()] = validation.FieldValidator(fe)
			}

			c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
			return

		}
		fmt.Println("err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	fmt.Println("Registered user", userinput)

	userEmail, err := services.RegisterUser(c.Request.Context(), userinput)

	if err != nil {
		fmt.Println("Err", err)
		if err.Error() == "username or email already exists" {
			c.JSON(http.StatusConflict, gin.H{"errors": map[string]string{"error": err.Error()}})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed due to an unknown server error."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully.", "userEmail": userEmail})
}

// LoginHandler handles the POST /api/auth/login request.
// RESPONSIBILITY: Validates credentials, calls the service to generate tokens, and sets the cookie.
func LoginHandler(c *gin.Context) {

	var loginInput models.LoginData

	if err := c.ShouldBindJSON(&loginInput); err != nil {

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

	accessToken, refreshToken, userRole, userName, userEmail, err := services.LoginUser(loginInput)

	if err != nil {

		fmt.Println("Error in login",err)
		if err.Error() == "your account has been blocked by admin" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "email not verified" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err.Error()=="invalid credentials"{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong. Please try again later."})
		return
	}

	c.SetCookie(cookieName, refreshToken, cookieMaxAge, "/", cookieDomain, false, true)
	c.JSON(200, gin.H{"accessToken": accessToken, "role": userRole, "name": userName, "email": userEmail})

}

// RefreshHandler handles the POST /api/auth/refresh request.
// RESPONSIBILITY: Reads the Refresh Token cookie, calls the service to rotate the tokens,
// and sends back a new pair. This endpoint is crucial for maintaining a logged-in session.
func RefreshHandler(c *gin.Context) {
	refreshTokenString, err := c.Cookie(cookieName)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
		return
	}

	accessToken, refreshToken, err := services.CheckRefreshToken(refreshTokenString)

	if err != nil {
		c.SetCookie("user_session", "", -1, "/", "localhost", false, true)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Session expired or revoked"})
		return
	}

	c.SetCookie(cookieName, refreshToken, cookieMaxAge, "/", cookieDomain, false, true)

	c.JSON(200, gin.H{"accessToken": accessToken})

}

// [LogoutHandler handles the POST /api/auth/logout request.
// RESPONSIBILITY: Clears the secure Refresh Token cookie to end the user's session.
func LogoutHandler(c *gin.Context) {
	userId := c.GetUint("userID")

	err := services.LogoutUser(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Session expired or revoked"})
		return
	}

	c.SetCookie("user_session", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully."})

}
