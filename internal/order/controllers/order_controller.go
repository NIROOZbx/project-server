package controllers

import (
	"fmt"
	"net/http"
	"strings"

	orderModel "github.com/NIROOZbx/project-server/internal/order/models"
	orderService "github.com/NIROOZbx/project-server/internal/order/services"
	"github.com/NIROOZbx/project-server/internal/shared/dtos"
	"github.com/NIROOZbx/project-server/internal/shared/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PlaceOrderHandler(c *gin.Context) {

	var orderData orderModel.PlaceOrderInput

	if err := c.ShouldBindJSON(&orderData); err != nil {
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			errorMessages := make(map[string]string)

			for _, fe := range validationErr {
				errorMessages[fe.Field()] = validation.FieldValidator(fe)
			}

			c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
			return

		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	userId := c.GetUint("userID")

	orderRef, err := orderService.PlaceOrderService(userId, orderData.AddressID, orderData.PaymentMethod)

	if err != nil {
		errMsg := err.Error()

		if errMsg == "address not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Selected address not found"})
			return
		}

		if errMsg == "cart is empty" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Your cart is empty"})
			return
		}

		if strings.HasPrefix(errMsg, "out of stock") {
			c.JSON(http.StatusConflict, gin.H{"error": errMsg})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to place order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":         "Order placed successfully",
		"order_reference": orderRef,
	})
}

func GetOrderHandler(c *gin.Context) {
	userID := c.GetUint("userID")

	orders, err := orderService.GetOrderService(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch your orders",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Orders fetched successfully",
		"orders":  orders,
	})
}

func UpdateOrderStatusHandler(c *gin.Context) {
	

	itemId := c.Param("itemID")

	var req orderModel.OrderStatus
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	err := orderService.UpdateOrderStatusService(req.Status, itemId)

	if err != nil {
		if err.Error() == "invalid order status" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "order not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if strings.Contains(err.Error(), "cannot update status") {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order status updated successfully"})
}

func GetAllOrderHandler(c *gin.Context) {

	var input dtos.PaginationInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	response, err := orderService.GetAllOrderService(input.Page, input.Limit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)

}

func CancelOrderHandler(c *gin.Context) {

	var input orderModel.OrderCancelReasons

	if err:=c.ShouldBindJSON(&input);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
        return
	}

	fmt.Printf("DEBUG: Received Reason: '%s'\n", input.CancelledReason)

	orderId := c.Param("id")

	userId := c.GetUint("userID")

	itemIdStr := c.Param("itemId")

	err := orderService.CancelOrderService(userId, orderId, itemIdStr,input.CancelledReason)

	if err != nil {
		errMsg := err.Error()

		if errMsg == "unauthorized: order does not belong to user" {
			c.JSON(http.StatusForbidden, gin.H{"error": errMsg})
			return
		}

		if err.Error() == "order not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if strings.Contains(errMsg, "cannot cancel") || strings.Contains(errMsg, "item is already") {
             c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
             return
        }

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order cancelled successfully"})
}

func GetDashboardStatsHandler(c *gin.Context) {
	stats, err := orderService.GetDashboardStatsService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stats"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"stats": stats,
	})
}
