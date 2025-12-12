package services

import (
	"errors"
	"strconv"
	"strings"

	wishlistModel "github.com/NIROOZbx/project-server/internal/wishlist/models"
	wishListRepo "github.com/NIROOZbx/project-server/internal/wishlist/repositories"
	"gorm.io/gorm"
)

func AddToWishlistService(userId uint, productId uint) error {

	err := wishListRepo.AddToWishlistInDB(userId, productId)

	if err != nil {
		errMsg := err.Error()

		if strings.Contains(errMsg, "duplicate key value") || strings.Contains(errMsg, "unique constraint") {
			return errors.New("product already in wishlist")
		}

		if strings.Contains(errMsg, "foreign key constraint") {
			return errors.New("product not found")
		}

		return err
	}

	return nil

}

func RemoveFromWishlistService(userId uint, productId string) error {

	convertedProdId, err := strconv.Atoi(productId)

	if err != nil {
		return errors.New("invalid product ID format")
	}
	err = wishListRepo.RemoveFromWishlistInDB(userId, uint(convertedProdId))

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("product not found in wishlist")
		}
		return err
	}

	return nil
}

func GetWishlistService(userId uint) ([]wishlistModel.WishlistItemDTO, error) {

	wishlistItems, err := wishListRepo.GetWishlistFromDB(userId)
	if err != nil {
		return nil, err
	}

	dtos := make([]wishlistModel.WishlistItemDTO, 0, len(wishlistItems))

	for _, item := range wishlistItems {
		p := item.Product

		dtos = append(dtos, wishlistModel.WishlistItemDTO{
			Id:       p.ID,
			Name:     p.Name,
			League:   p.League,
			Price:    p.Price,
			ImageURL: p.Image,
		})

	}

	return dtos, nil

}
