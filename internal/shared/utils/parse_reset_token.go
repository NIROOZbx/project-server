package utils

import (
	"errors"

	"github.com/NIROOZbx/project-server/internal/shared/config"
	"github.com/golang-jwt/jwt/v5"
)

func ParseResetToken(resetToken string) (uint, error) {

	resetSecret := config.GetConfig().ResetSecret

	token, err := jwt.Parse(resetToken, func(t *jwt.Token) (any, error) {
		return []byte(resetSecret), nil
	})

	if err != nil || !token.Valid {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return 0, errors.New("invalid token claims")
    }
	if purpose, ok := claims["purpose"].(string); !ok || purpose != "reset-password" {
        return 0, errors.New("invalid token type")
    }

	sub, ok := claims["id"].(float64) 
    if !ok {
        return 0, errors.New("token does not contain valid user ID")
    }

	return uint(sub),nil

}
