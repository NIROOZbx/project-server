package services

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	cartRepo "github.com/NIROOZbx/project-server/internal/cart/repositories"
	"github.com/NIROOZbx/project-server/internal/global"
	orderModel "github.com/NIROOZbx/project-server/internal/order/models"
	orderRepo "github.com/NIROOZbx/project-server/internal/order/repositories"
	"gorm.io/gorm"
)

func GenerateOrderReference() string {

	datePart := time.Now().Format("20060102")

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	randomPart := rng.Intn(10000)

	return fmt.Sprintf("ORD-%s-%04d", datePart, randomPart)
}

func PlaceOrderService(userId uint, addressId uint, payMode string) (string, error) {

	err := orderRepo.FindAddressById(userId, addressId)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("address not found")
		}

		return "", err
	}

	cartItems, err := cartRepo.GetCartFromDB(userId)

	if err != nil {
		return "", errors.New("can't find user")
	}

	if len(cartItems) == 0 {
		return "", errors.New("cart is empty")
	}

	var totalPrice float64

	for _, item := range cartItems {
		if item.Quantity > item.Product.Stock {
			return "", errors.New("only this much product left")
		}
		totalPrice += float64(item.Quantity) * item.Product.Price
	}

	refID := GenerateOrderReference()
	newOrder := &orderModel.Order{
		UserID:        userId,
		AddressID:     addressId,
		TotalPrice:    totalPrice,
		PaymentMethod: payMode,
		Status:        "pending",
		Reference:     refID,
	}

	newPayment:=&orderModel.Payment{
		UserID: userId,
		Amount: totalPrice,
		PaymentMethod: payMode,
		PaymentStatus: "pending",
	}

	err = orderRepo.PlaceOrderInDB(newOrder, cartItems,newPayment)

	if err != nil {
		return "", err
	}

	if global.NotifService!=nil{
		title := "Order Confirmed! 🎉"
        body := "Your order #" + refID + " has been placed successfully. Wait for more updates"
		global.NotifService.SendNotificationService(userId,title,body)
	}



	return refID, nil

}

func GetOrderService(userId uint) ([]orderModel.OrderResponse, error) {

	orders, err := orderRepo.GetOrderFromDB(userId)

	if err != nil {
		return nil, err
	}
	orderResponses := make([]orderModel.OrderResponse,0, len(orders))


	for _, order := range orders {
		var itemsDTO []orderModel.OrderItemDTO

		for _, item := range order.Items {
			
			itemsDTO = append(itemsDTO, orderModel.OrderItemDTO{
				ItemID: item.ID,
				ProductID:       item.ProductID,
				ProductName:     item.ProductName,
				ProductImageURL: item.ImageURL,
				PriceAtPurchase: item.PriceAtPurchase,
				Quantity:        item.Quantity,
				ItemStatus:item.ItemStatus,
			})
		}

			orderResponses = append(orderResponses, orderModel.OrderResponse{
				OrderID:   order.ID,
				OrderDate: order.OrderDate,
				Status:    order.Status,
				Reference: order.Reference,
				Address: orderModel.AddressDTO{
					Name:          order.Address.Name,
					Phone:         order.Address.Phone,
					StreetAddress: order.Address.StreetAddress,
					City:          order.Address.City,
					State:         order.Address.State,
					ZipCode:       order.Address.ZipCode,
					Country:       order.Address.Country,
				},
				PaymentMethod: order.PaymentMethod,
				Items:         itemsDTO,
				TotalPrice:    order.TotalPrice,
			})
		
	}
	

	return orderResponses, nil

}



func CancelOrderService(userId uint,orderIdStr string,itemIdStr,CancelledReason string)error{

	orderID,err:=strconv.Atoi(orderIdStr)

	if err != nil {
		return errors.New("invalid order ID format")
	}
	itemID,err:=strconv.Atoi(itemIdStr)

	if err != nil {
		return errors.New("invalid order ID format")
	}




	err=orderRepo.CancelOrderInDB(uint(orderID),userId,uint(itemID),CancelledReason)

	if err != nil {
		if err==gorm.ErrRecordNotFound{ 
		return errors.New("order not found")
		}
		return err
	}

	return nil

}
