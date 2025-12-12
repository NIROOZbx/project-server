package respositories

import (

	authModel "github.com/NIROOZbx/project-server/internal/auth/models"
	"github.com/NIROOZbx/project-server/internal/shared/database"
	"gorm.io/gorm"
)

func GetUserFromDB(userId uint) (*authModel.User, error) {
	var db = database.DB()
	var userProfile authModel.User

	res := db.First(&userProfile, userId)

	return &userProfile, res.Error
}

func ChangeNameInDB(userId uint, name string) error {
	var db = database.DB()

	res := db.Model(&authModel.User{}).Where("id = ?", userId).Update("name", name)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}

func GetUsersWithPagination(limit int, offset int) ([]authModel.User, error) {

	var db = database.DB()

	var allUser []authModel.User

	res := db.Where("role = ?", "user").Limit(limit).Offset(offset).Find(&allUser)

	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return allUser, res.Error

}

func CountUsers() (int64, error) {
	db := database.DB()
	var count int64
	err := db.Model(&authModel.User{}).Where("role = ?", "user").Count(&count).Error
	return count, err
}


func BlockUserInDB(userID uint,blockUser bool) error{

	var db=database.DB()

	res:=db.Model(&authModel.User{}).Where("id = ?", userID).Update("is_blocked",blockUser)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil



}


func UploadProfileImageInDB(imageURL string,userID uint)error{

	var db=database.DB()

	result := db.Model(&authModel.User{}).Where("id = ?", userID).Update("profile_image", imageURL)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil

}