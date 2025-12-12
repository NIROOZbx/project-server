package services

import (
	"errors"
	"time"

	"github.com/NIROOZbx/project-server/internal/auth/repositories"
	"github.com/NIROOZbx/project-server/internal/shared/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userID uint, Role string, expDuration time.Duration, tokenVersion int) (string, error) {

	accessSecret := config.GetConfig().AccessSecret

	claims := jwt.MapClaims{
		"userID":  userID,
		"Role":    Role,
		"version": tokenVersion,
		"exp":     jwt.NewNumericDate(time.Now().Add(expDuration)),
		"iat":     jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(accessSecret))

	if err != nil {
		return "", err

	}

	return tokenString, nil

}


func GenerateResetToken(email string,userID uint) (string, error) {
	resetSecret := config.GetConfig().ResetSecret
	claims := jwt.MapClaims{
		"id":userID,
		"email":   email,
		"purpose": "reset-password",
		"exp":     jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(resetSecret))

	if err != nil {

		return "", err

	}

	return tokenString, nil
}


func GenerateRefreshToken(userID uint, Role string, expDuration time.Duration, tokenVersion int) (string, error) {
	refreshSecret := config.GetConfig().RefreshSecret

	claims := jwt.MapClaims{
		"userID":  userID,
		"Role":    Role,
		"version": tokenVersion,
		"exp":     jwt.NewNumericDate(time.Now().Add(expDuration)),
		"iat":     jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(refreshSecret))

	if err != nil {
		return "", err

	}

	return tokenString, nil
}

func CheckRefreshToken(refreshToken string) (string, string, error) {

	refreshSecret := config.GetConfig().RefreshSecret
	accessSecret := config.GetConfig().AccessSecret

	token, err := jwt.Parse(refreshToken, func(t *jwt.Token) (any, error) {
		return []byte(refreshSecret), nil
	})

	if err != nil || !token.Valid {
		return "", "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return "", "", errors.New("invalid token claims")
	}

	userID := uint(claims["userID"].(float64))
	Role := claims["Role"].(string)
	version := int(claims["version"].(float64))

	user, DbErr := repositories.FindUserById(userID)

	if DbErr != nil {
		return "", "", errors.New("user session revoked or deleted")
	}

	if user.TokenVersion != version {
		return "", "", errors.New("invalid token version")
	}

	newClaims := jwt.MapClaims{
		"userID":  userID,
		"Role":    Role,
		"version": version,
		"exp":     jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
		"iat":     jwt.NewNumericDate(time.Now()),
	}

	newRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)

	fullRefreshToken, err := newRefreshToken.SignedString([]byte(refreshSecret))

	if err != nil {
		return "", "", err
	}

	accessClaims := jwt.MapClaims{
		"userID":  userID,
		"Role":    Role,
		"version": version,
		"exp":     jwt.NewNumericDate(time.Now().Add(20 * time.Minute)),
		"iat":     jwt.NewNumericDate(time.Now()),
	}

	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	fullAccessToken, err := newAccessToken.SignedString([]byte(accessSecret))

	if err != nil {
		return "", "", err
	}

	return fullAccessToken, fullRefreshToken, nil

}
