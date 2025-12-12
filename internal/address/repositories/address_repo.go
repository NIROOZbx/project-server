package repositories

import (
	"errors"
	"strings"

	addressModel "github.com/NIROOZbx/project-server/internal/address/models"
	"github.com/NIROOZbx/project-server/internal/shared/database"
	"gorm.io/gorm"
)

func CreateAddress(address *addressModel.Address) error {
	db := database.DB()

	err := db.Create(address).Error

	if err != nil {

		if strings.Contains(err.Error(), "foreign key constraint") {
			return errors.New("user not found")
		}

	}
	return err
}

func CountAddresses(userId uint) (int64, error) {

	var db = database.DB()

	var count int64

	res := db.Model(&addressModel.Address{}).Where("user_id = ?", userId).Count(&count).Error

	return count, res

}

func GetAddressFromDB(userId uint) ([]addressModel.Address, error) {

	db := database.DB()
	var addresses []addressModel.Address

	result := db.Where("user_id = ?", userId).Find(&addresses)

	return addresses, result.Error

}

func RemoveAddressFromDB(userId uint, addressId uint) error {

	db := database.DB()

	result := db.Where("user_id = ? AND id = ?", userId, addressId).Delete(&addressModel.Address{})

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error

}

func SetDefaultAddressInDB(userId uint, addressId uint) error {

	db := database.DB()

	res := db.Model(&addressModel.Address{}).
		Where("user_id = ?", userId).
		Update("is_default", gorm.Expr("id = ?", addressId))

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return res.Error

}
