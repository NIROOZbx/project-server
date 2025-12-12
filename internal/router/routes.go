package routes

import (
	addressController "github.com/NIROOZbx/project-server/internal/address/controllers"
	authController "github.com/NIROOZbx/project-server/internal/auth/controllers"
	cartController "github.com/NIROOZbx/project-server/internal/cart/controller"
	notifController "github.com/NIROOZbx/project-server/internal/notification/controllers"
	orderContoller "github.com/NIROOZbx/project-server/internal/order/controllers"
	productController "github.com/NIROOZbx/project-server/internal/products/controllers"
	reviewController "github.com/NIROOZbx/project-server/internal/reviews/controllers"
	"github.com/NIROOZbx/project-server/internal/shared/middlewares"
	userController "github.com/NIROOZbx/project-server/internal/user/controllers"
	wishlistController "github.com/NIROOZbx/project-server/internal/wishlist/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(router *gin.Engine,notifHandler *notifController.NotificationHandler) {

	//AUTH ROUTES
	auth := router.Group("/auth")
	{
		auth.POST("/signup", authController.RegisterHandler)
		auth.POST("/login", authController.LoginHandler)
		auth.POST("/refresh", authController.RefreshHandler)
		auth.POST("/logout", middlewares.AuthMiddleware(), authController.LogoutHandler)
		auth.POST("/verify-otp", authController.VerifyOTPHandler)
		auth.POST("/resend-otp", authController.ResendOTPHandler)

		auth.POST("/forgot-password", authController.ForgotPasswordHandler)
		auth.POST("/verify-reset-otp", authController.VerifyResetOTPHandler)
		auth.POST("/reset-password",middlewares.RequireResetToken(), authController.ResetPasswordHandler)

	}

	//PRODUCT ROUTES
	productRoutes := router.Group("/public")
	{
		productRoutes.GET("/products", productController.ListProductsHandler)
		productRoutes.GET("/products/:id", productController.ShowSingleProductHandler)
		productRoutes.GET("/home-products", productController.GetHomeProducts)
		productRoutes.GET("/:id/reviews", reviewController.GetProductReviewsHandler)
	}

	//USER ROUTES
	api := router.Group("/api", middlewares.AuthMiddleware(), middlewares.RoleMiddleware("user"))
	{
		userRoutes := api.Group("/user")
		{
			userRoutes.GET("/profile", userController.ShowUserProfile)
			userRoutes.PATCH("/update-name", userController.UpdateUserName)
			userRoutes.PATCH("/update-profile", userController.UpdateUserProfileImage)
			userRoutes.POST("/change-password", userController.UpdateUserPassword)
		}

		reviewRoutes := api.Group("/review")
		{
			reviewRoutes.POST("/:id", reviewController.AddReviewHandler)
		}

		//CART ROUTES
		cartRoutes := api.Group("/cart")
		{
			cartRoutes.GET("/", cartController.GetCartHandler)
			cartRoutes.POST("/", cartController.AddToCartHandler)
			cartRoutes.PATCH("/", cartController.UpdateCartQuantityHandler)
			cartRoutes.DELETE("/:id", cartController.RemoveFromCartHandler)
		}

		//WISHLIST ROUTES

		wishlistRoutes := api.Group("/wishlist")
		{
			wishlistRoutes.GET("/", wishlistController.GetWishlistHandler)
			wishlistRoutes.POST("/", wishlistController.AddToWishlistHandler)
			wishlistRoutes.DELETE("/:id", wishlistController.RemoveWishlistHandler)
		}

		//ADDRESS ROUTES

		addressGroup := api.Group("/addresses")

		{
			addressGroup.POST("/", addressController.AddAddressHandler)
			addressGroup.GET("/", addressController.GetAddressesHandler)
			addressGroup.DELETE("/:id", addressController.DeleteAddressHandler)
			addressGroup.PATCH("/:id", addressController.SetDefaultAddressHandler)
		}


		//ORDER ROUTES

		orderGroup := api.Group("/orders")
		{
			orderGroup.POST("/", orderContoller.PlaceOrderHandler)
			orderGroup.GET("/", orderContoller.GetOrderHandler)
			orderGroup.PATCH("/:id/cancel-order/:itemId", orderContoller.CancelOrderHandler)
		}


		//NOTIFICATION ROUTES

		notifRoutes := api.Group("/notifications")

	{
		notifRoutes.GET("/", notifHandler.GetNotifications)
		notifRoutes.PATCH("/:id/read", notifHandler.MarkRead)
		notifRoutes.PATCH("/read-all", notifHandler.MarkAllRead)
	}

		adminRoutes := router.Group("/admin", middlewares.AuthMiddleware(), middlewares.RoleMiddleware("admin"))
		{
			//ADMIN USER MANAGEMENT ROUTES
			adminRoutes.GET("/users", userController.GetAlluserHandler)
			adminRoutes.PATCH("/users/:id", userController.BlockUserHandler)

			//ADMIN PRODUCT MANAGEMENT ROUTES
			adminRoutes.POST("/products", productController.AddProductHandler)
			adminRoutes.DELETE("/products/:id", productController.DeleteProductHandler)
			adminRoutes.GET("/products", productController.GetAllProductsHandler)
			adminRoutes.PATCH("/products/:id", productController.UpdateProductHandler)

			//ADMIN ORDER MANAGEMENT ROUTES
			adminRoutes.PATCH("/orders/:itemID", orderContoller.UpdateOrderStatusHandler)
			adminRoutes.GET("/orders", orderContoller.GetAllOrderHandler)
			adminRoutes.GET("/dashboard-stats", orderContoller.GetDashboardStatsHandler)

			adminRoutes.POST("/notifications", notifHandler.AddNotifcationHandler)
		}

	}

}
