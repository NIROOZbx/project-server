package services

import (
	"context"
	"errors"

	"github.com/NIROOZbx/project-server/internal/auth/models"

	authRepo "github.com/NIROOZbx/project-server/internal/auth/repositories"
	"github.com/NIROOZbx/project-server/internal/shared/security"

)

func RegisterUser(ctx context.Context, userinput models.SignInData) (string, error) {

	if authRepo.FindDuplicateUser(userinput.Email,userinput.Name){
		return "", errors.New("username or email already exists")
	}
	
	hashedPassword, hashErr := security.HashPassword(userinput.Password)

	if hashErr != nil {
		return "", hashErr
	}

	var newUser = &models.User{
		Name:     userinput.Name,
		Password: hashedPassword,
		Email:    userinput.Email,
	}

	if err := authRepo.CreateUser(newUser); err != nil {
		return "", err
	}
	err:= SendOTP(ctx, newUser.Email)

	if err != nil {
		return newUser.Email, err
	}

	return newUser.Email, nil

}
