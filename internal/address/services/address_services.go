package services

import (
	"errors"

	"strconv"

	addressModel "github.com/NIROOZbx/project-server/internal/address/models"
	addressRepo "github.com/NIROOZbx/project-server/internal/address/repositories"
	"gorm.io/gorm"
)

func AddAddressService(userId uint, addressInput addressModel.AddressInput) error {

	count, err := addressRepo.CountAddresses(userId)
	if err != nil {
		return err
	}
	if count >= 5 {
		return errors.New("maximum address limit reached")
	}

	if count == 0 {
		addressInput.IsDefault = true
	}

	newAddress := &addressModel.Address{
		UserID:        userId,
		Name:          addressInput.Name,
		Phone:         addressInput.Phone,
		StreetAddress: addressInput.StreetAddress,
		City:          addressInput.City,
		State:         addressInput.State,
		ZipCode:       addressInput.ZipCode,
		Country:       addressInput.Country,
		IsDefault:     addressInput.IsDefault,
	}

	return addressRepo.CreateAddress(newAddress)

}
func GetAddressService(userId uint) ([]addressModel.AddressResponseDTO, error) {

	addresses, err := addressRepo.GetAddressFromDB(userId)

	if err != nil {
		return nil, err
	}
	responseList := make([]addressModel.AddressResponseDTO, 0, len(addresses))

	for _, addr := range addresses {
		responseList = append(responseList, addressModel.AddressResponseDTO{
			ID:            addr.ID,
			Name:          addr.Name,
			Phone:         addr.Phone,
			StreetAddress: addr.StreetAddress,
			City:          addr.City,
			State:         addr.State,
			ZipCode:       addr.ZipCode,
			Country:       addr.Country,
			IsDefault:     addr.IsDefault,
		})
	}

	return responseList, nil

}

func DeleteAddressService(userId uint, addressId string) error {
	convertedAddressId, err := strconv.Atoi(addressId)

	if err != nil {
		return errors.New("invalid address ID format")
	}

	err = addressRepo.RemoveAddressFromDB(userId, uint(convertedAddressId))

	if err != nil {

		if err == gorm.ErrRecordNotFound {
			return errors.New("address not found or access denied")
		}

		return err
	}

	return nil

}

func UpdateDefaultAddressService(userId uint, addressId string) error {

	convertedAddressId, err := strconv.Atoi(addressId)

	if err != nil {
		return errors.New("invalid address ID format")
	}

	err = addressRepo.SetDefaultAddressInDB(userId, uint(convertedAddressId))

	if err != nil {

		if err == gorm.ErrRecordNotFound {
			return errors.New("address not found or access denied")
		}

		return err
	}

	return nil

}
