package services

import (
	"errors"
	"strconv"
	"strings"

	"github.com/NIROOZbx/project-server/internal/global"
	orderModel "github.com/NIROOZbx/project-server/internal/order/models"
	orderRepo "github.com/NIROOZbx/project-server/internal/order/repositories"
	"github.com/NIROOZbx/project-server/internal/shared/dtos"
	"gorm.io/gorm"
)

var validWords = map[string]bool{
	"shipped":   true,
	"delivered": true,
	"pending":   true,
	"cancelled": true,
}

func UpdateOrderStatusService(orderStatus, itemIdStr string) error {

	itemID, err := strconv.Atoi(itemIdStr)
	if err != nil {
		return errors.New("invalid query format")
	}

	if !validWords[orderStatus] {
		return errors.New("invalid order status")
	}

	err = orderRepo.UpdateOrderStatusInDB(orderStatus, uint(itemID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("order not found")
		}
		if strings.Contains(err.Error(), "cannot update status") {
			return err
		}
		return err
	}

	if global.NotifService != nil {
		userID, productName, err := orderRepo.GetUserIDByItemID(uint(itemID))

		title, body := getNotificationMessage(orderStatus, productName)

		if err == nil {
			global.NotifService.SendNotificationService(userID,
				title,
				body,
			)
		}
	}

	return nil

}

func getNotificationMessage(status, productName string) (string, string) {
	switch status {
	case "shipped":
		return "Order Shipped! 🚚", "Good news! Your " + productName + " is on its way."
	case "delivered":
		return "Order Delivered 🎉", "Your " + productName + " has been delivered. Enjoy!"
	case "cancelled":
		return "Order Cancelled ❌", "Your order for " + productName + " has been cancelled."
	default:
		return "Order Update", "Your item status has been updated to: " + status
	}
}

func GetAllOrderService(page, limit int) (*dtos.PaginationResponse, error) {

	offset := (page - 1) * limit

	totalCount, err := orderRepo.CountAllOrderItems()
	if err != nil {
		return nil, err
	}

	items, err := orderRepo.GetAllOrderItemsFromDB(limit, offset)

	if err != nil {
		return nil, err
	}

	orderResponses := make([]orderModel.AdminOrderDTO, 0, len(items))

	for _, item := range items {

		orderResponses = append(orderResponses, orderModel.AdminOrderDTO{
			OrderID:       item.OrderID,
			ItemID:        item.ID,
			ItemStatus:    item.ItemStatus,
			OrderDate:     item.Order.OrderDate,
			CustomerName:  item.Order.User.Name,
			CustomerEmail: item.Order.User.Email,
			ProductID:     item.ProductID,
			ProductName:   item.ProductName,
			ProductImage:  item.ImageURL,
			ProductPrice:  item.PriceAtPurchase,
			ProductLeague: item.Product.League,
			Quantity:      item.Quantity,
			TotalPrice:    item.Order.TotalPrice,
		})

	}

	totalPages := (totalCount + int64(limit) - 1) / int64(limit)

	return &dtos.PaginationResponse{
		Data:       orderResponses,
		TotalCount: totalCount,
		Page:       page,
		TotalPages: totalPages,
	}, nil
}

func GetDashboardStatsService() (*orderModel.OrderStats, error) {
	return orderRepo.GetOrderStatsFromDB()
}
