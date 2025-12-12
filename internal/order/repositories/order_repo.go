package repositories

import (
	"errors"
	
	"strings"

	addressModel "github.com/NIROOZbx/project-server/internal/address/models"
	authModel "github.com/NIROOZbx/project-server/internal/auth/models"
	cartModel "github.com/NIROOZbx/project-server/internal/cart/models"
	orderModel "github.com/NIROOZbx/project-server/internal/order/models"
	productModel "github.com/NIROOZbx/project-server/internal/products/models"
	"github.com/NIROOZbx/project-server/internal/shared/database"
	"gorm.io/gorm"
)

func PlaceOrderInDB(order *orderModel.Order, cartItems []cartModel.Cart, paymentTable *orderModel.Payment) error {

	db := database.DB()

	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(paymentTable).Error; err != nil {
		tx.Rollback()
		return err

	}
	for _, cartItem := range cartItems {

		orderItem := orderModel.OrderItem{
			OrderID:         order.ID,
			ProductID:       cartItem.ProductID,
			ProductName:     cartItem.Product.Name,
			PriceAtPurchase: cartItem.Product.Price,
			Quantity:        cartItem.Quantity,
			ImageURL:        cartItem.Product.Image,
			ItemStatus:      "active",
		}

		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			return err
		}

		result := tx.Model(&productModel.Product{}).
			Where("id = ? AND stock >= ?", cartItem.ProductID, cartItem.Quantity).
			Update("stock", gorm.Expr("stock - ?", cartItem.Quantity))

		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}

		if result.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("product out of stock: " + cartItem.Product.Name)
		}
	}

	if err := tx.Where("user_id = ?", order.UserID).Delete(&cartModel.Cart{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func GetOrderFromDB(userId uint) ([]orderModel.Order, error) {
	db := database.DB()
	var order []orderModel.Order
	result := db.Preload("Items").Preload("Address").Where("user_id = ?", userId).Order("created_at DESC").Find(&order)

	if result.Error != nil {
		return nil, result.Error
	}

	return order, nil

}

func FindAddressById(userId uint, addressId uint) error {

	var db = database.DB()

	res := db.First(&addressModel.Address{}, "id = ? AND user_id = ?", addressId, userId)

	return res.Error

}

func UpdateOrderStatusInDB(newStatus string, itemID uint) error {

	var db = database.DB()

	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var item orderModel.OrderItem
	if err := tx.First(&item, itemID).Error; err != nil {
		tx.Rollback()
		return err
	}
	oldStatus := strings.ToLower(item.ItemStatus)

	if oldStatus == "cancelled" || oldStatus=="delivered"{
        tx.Rollback()
        return errors.New("cannot update status: item is already "+oldStatus)
    }

	if (newStatus == "cancelled") && (oldStatus != "cancelled") {
		if err := tx.Model(&productModel.Product{}).
			Where("id = ?", item.ProductID).
			Update("stock", gorm.Expr("stock + ?", item.Quantity)).Error; err != nil {
			tx.Rollback()
			return err
		}

        refundAmount := item.PriceAtPurchase * float64(item.Quantity)

        if err := tx.Model(&orderModel.Order{}).
            Where("id = ?", item.OrderID).
            Update("total_price", gorm.Expr("total_price - ?", refundAmount)).Error; err != nil {
            tx.Rollback()
            return err
        }
	}

	if err := tx.Model(&item).Update("item_status", newStatus).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := checkAndUpdateParentOrder(tx, item.OrderID); err != nil {
         tx.Rollback()
         return err
    }

	return tx.Commit().Error

}



func GetAllOrderItemsFromDB(limit, offset int) ([]orderModel.OrderItem, error) {
    db := database.DB()
    var items []orderModel.OrderItem

    err := db.Model(&orderModel.OrderItem{}).
        Preload("Order").
		Preload("Product").
        Preload("Order.User").
        Joins("JOIN orders ON orders.id = order_items.order_id").
        Order("orders.created_at DESC").
        Limit(limit).  
        Offset(offset).
        Find(&items).Error

    return items, err
}

func CountAllOrderItems() (int64, error) {
	var count int64
	err := database.DB().Model(&orderModel.OrderItem{}).Count(&count).Error
	return count, err
}

func CancelOrderInDB(orderId uint, userID uint, itemId uint,cancelledReason string) error {

	var db = database.DB()

	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var order orderModel.Order

	if err := tx.Preload("Items").First(&order, orderId).Error; err != nil {
		tx.Rollback()
		return err
	}

	if order.UserID != userID {
		tx.Rollback()
		return errors.New("unauthorized: order does not belong to user")
	}

	if order.Status != "pending" && order.Status != "shipped" {
		tx.Rollback()
		return errors.New("cannot cancel order: status is " + order.Status)
	}

	var item orderModel.OrderItem

	if err := tx.First(&item, "id = ? AND order_id = ?", itemId, orderId).Error; err != nil {
		tx.Rollback()
		return errors.New("order item not found")
	}

	if item.ItemStatus == "cancelled" {
		tx.Rollback()
		return errors.New("item already cancelled")
	}

	if err := tx.Model(&item).Updates(map[string]interface{}{
        "item_status":      "cancelled",
        "cancelled_reason": cancelledReason,
    }).Error; err != nil {
        tx.Rollback()
        return err
    }

	if err := tx.Model(&productModel.Product{}).Where("id = ?", item.ProductID).
		Update("stock", gorm.Expr("stock + ?", item.Quantity)).Error; err != nil {

		tx.Rollback()
		return err
	}

	if err := tx.Model(&order).Update("total_price", gorm.Expr("total_price - ?", item.PriceAtPurchase)).Error; err != nil {

		tx.Rollback()
		return err
	}

	var cancelCount int64

	err := tx.Model(&orderModel.OrderItem{}).
		Where("order_id = ? AND item_status = ?", order.ID, "cancelled").
		Count(&cancelCount).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	if cancelCount == int64(len(order.Items)) {
		if err := tx.Model(&order).Update("status", "cancelled").Error; err != nil {
			tx.Rollback()
			return err
		}

	}

	return tx.Commit().Error

}
func checkAndUpdateParentOrder(tx *gorm.DB, orderID uint) error {
    var order orderModel.Order
    
    
    if err := tx.Preload("Items").First(&order, orderID).Error; err != nil {
        return err
    }

    totalItems := int64(len(order.Items))
    var cancelledCount int64
    var deliveredCount int64

    for _, item := range order.Items {
      
        status := strings.ToLower(item.ItemStatus)
        if status == "cancelled" {
            cancelledCount++
        } else if status == "delivered" {
            deliveredCount++
        }
    }

    
    if cancelledCount == totalItems {
        return tx.Model(&order).Update("status", "cancelled").Error
    }

    if deliveredCount==totalItems {
        return tx.Model(&order).Update("status", "delivered").Error
    }

    return nil
}


func GetOrderStatsFromDB() (*orderModel.OrderStats, error) {
    db := database.DB()

	stats := &orderModel.OrderStats{}
    

    
    var results []orderModel.Result

    err := db.Model(&orderModel.OrderItem{}).
       Select("item_status, count(*) as count, COALESCE(SUM(price_at_purchase * quantity), 0) as total_amount").
        Group("item_status").
        Scan(&results).Error

		if err := db.Model(&productModel.Product{}).Count(&stats.TotalProducts).Error; err != nil {
        return nil, err
    }
	if err := db.Model(&authModel.User{}).Where("role = ?","user").Count(&stats.TotalUsers).Error; err != nil {
        return nil, err
    }

    if err != nil {
        return nil, err
    }

    for _, r := range results {
        stats.TotalItems += r.Count
        
        switch r.ItemStatus {
        case "active", "processing": 
            stats.Active += r.Count
        case "shipped":
            stats.Shipped = r.Count
        case "delivered":
            stats.Delivered = r.Count
			stats.TotalRevenue = r.TotalPrice
			stats.EstimatedProfit=r.TotalPrice*0.3
        case "cancelled":
            stats.Cancelled = r.Count
        }
    }

    return stats, nil
}

func GetUserIDByItemID(itemID uint) (uint,string, error) {
    var db = database.DB()
    var result struct {
        UserID      uint
        ProductName string
    }

   
    err := db.Table("order_items").
        Select("orders.user_id, order_items.product_name").
        Joins("JOIN orders ON orders.id = order_items.order_id").
        Where("order_items.id = ?", itemID).
        Scan(&result).Error

    if err != nil {
        return 0,"", err
    }

    return result.UserID,result.ProductName, nil
}
