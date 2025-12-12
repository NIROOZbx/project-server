package services

import (
	"errors"
	"fmt"

	"strconv"

	"strings"

	cartModel "github.com/NIROOZbx/project-server/internal/cart/models"
	cartRepo "github.com/NIROOZbx/project-server/internal/cart/repositories"
	"gorm.io/gorm"
)

func AddToCartService(userId uint, productId uint) error {

	err := cartRepo.AddToCartInDB(userId, productId)

	if err != nil {
		errMsg := err.Error()
		if strings.Contains(err.Error(), "duplicate key value") {
			return errors.New("product already exists in the cart")
		}
		if strings.Contains(errMsg, "foreign key constraint") {
			return errors.New("product does not exist")
		}
		return err
	}

	return nil

}

func UpdateCartQuantityService(userId uint, productId uint, quantity int) error {

	product,err:=cartRepo.FindProductById(productId)

	if err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("product not found")
        }
		return err
	}

	fmt.Println("Current user quantity",quantity)
	
	if quantity>product.Stock{
		return errors.New("no stock left")
	}

	if quantity>=3{
		return errors.New("maximum quantity reached")
	}

	err = cartRepo.UpdateCartQuantityInDB(userId, productId, quantity)

	if err == gorm.ErrRecordNotFound {
		return errors.New("item not found in cart")
	}

	return err

}

func RemoveFromCartService(userId uint, productId string) error {

	convertedProdId,err:=strconv.Atoi(productId)

	if err != nil {
		return errors.New("invalid product ID format")
	}

	err = cartRepo.RemoveFromCartInDB(userId, uint(convertedProdId))

	if err != nil {
		return err
	}

	return nil

}

func GetCartService(userId uint) (*cartModel.CartResponseDTO, error) {



	userCart, err := cartRepo.GetCartFromDB(userId)

	if err != nil {
		return nil, err
	}
	cartItemsDTO:=make([]cartModel.CartItemDTO,0,len(userCart))

	var totalItems int
	var totalPrice float64

	for _,item :=range userCart{

		

		totalItems+=item.Quantity
		totalPrice+=float64(item.Quantity) * item.Product.Price

		cartItemsDTO = append(cartItemsDTO, cartModel.CartItemDTO{
			ID: item.ProductID,
			Name: item.Product.Name,
			Price: item.Product.Price,
			League: item.Product.League,
			ImageURL: item.Product.Image,
			Quantity: item.Quantity,
			SubTotal: float64(item.Quantity)*item.Product.Price,

		})
		
	}

	cartResponseDTO:=&cartModel.CartResponseDTO{
		Items: cartItemsDTO,
		TotalItems: totalItems,
		TotalPrice: totalPrice,
	}



	return cartResponseDTO,nil

}
