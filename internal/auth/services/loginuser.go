package services

import (
	"errors"
	"time"

	"github.com/NIROOZbx/project-server/internal/auth/models"
	"github.com/NIROOZbx/project-server/internal/auth/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginUser(LoginData models.LoginData) (string, string,string,string,string,error) {

	userData, err := repositories.FindUserByEmail(LoginData.Email)

	if err == gorm.ErrRecordNotFound {
		return "", "","","","", errors.New("invalid credentials")
	}
	
	if !userData.IsVerified {
		return "", "","","","", errors.New("email not verified")
	}
	

	if userData.IsBlocked {
		
		return "", "","","","", errors.New("your account has been blocked by admin")
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(LoginData.Password))

	if passErr != nil {
		return "", "","","","", errors.New("invalid credentials")
	}

	accessToken, err := GenerateAccessToken(userData.ID, userData.Role, 20*time.Minute, userData.TokenVersion)
	if err != nil {

		return "", "","","","",err

	}
	refreshToken, err := GenerateRefreshToken(userData.ID, userData.Role, 24*7*time.Hour, userData.TokenVersion)
	if err != nil {

		return "", "","","","", err

	}

	return accessToken, refreshToken, userData.Role,userData.Name,userData.Email,nil

}
