package services

import (
	"context"
	"errors"
	"mime/multipart"

	authRepo "github.com/NIROOZbx/project-server/internal/auth/repositories"

	fileupload "github.com/NIROOZbx/project-server/internal/shared/fileUpload"
	"github.com/NIROOZbx/project-server/internal/shared/security"
	userModel "github.com/NIROOZbx/project-server/internal/user/models"
	userRepo "github.com/NIROOZbx/project-server/internal/user/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetUserProfile(userId uint) (*userModel.UserProfile, error) {

	user, err := userRepo.GetUserFromDB(userId)

	if err != nil {
		return nil, err
	}

	profile := &userModel.UserProfile{
		Name:       user.Name,
		Email:      user.Email,
		IsVerified: user.IsVerified,
		ProfileImage: user.ProfileImage,
	}

	return profile, nil

}

func UpdateUserName(userId uint, newName string) error {

	err := userRepo.ChangeNameInDB(userId, newName)

	if err == gorm.ErrRecordNotFound {
		return errors.New("no user found")
	}

	return err

}

func UpdatePassword(userId uint, newPassword string, oldPassword string) error {

	user, err := authRepo.FindUserById(userId)

	if err == gorm.ErrRecordNotFound {
		return errors.New("user not found")
	}
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))

	if err != nil {
		return errors.New("the old password you entered is incorrect")
	}

	if oldPassword == newPassword {
		return errors.New("cannot be same as old password")
	}

	hashedPassword, err := security.HashPassword(newPassword)

	if err != nil {
		return err
	}

	err = authRepo.UpdatePassword(userId, hashedPassword)

	if err != nil {

		return err
	}

	return nil
}


func UploadProfileImage(ctx context.Context,userID uint,imageData *multipart.FileHeader)(string,error){

	
	imageURL, err := fileupload.UploadFileToCloudinary(ctx, imageData)
	if err != nil {
		return "", errors.New("image upload failed")
	}

	err=userRepo.UploadProfileImageInDB(imageURL, userID)

	if err != nil {
		return "", err
	}

	return imageURL, nil


}