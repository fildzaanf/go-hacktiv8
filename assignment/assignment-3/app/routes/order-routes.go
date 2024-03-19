package routes

import (
	"assignment-2/middlewares"
	"assignment-2/modules/order/handler"
	"assignment-2/modules/order/repository"
	"assignment-2/modules/order/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func OrderRoutes(r *gin.RouterGroup, db *gorm.DB) {
	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepository)
	orderHandler := handler.NewOrderHandler(orderService)

	order := r.Group("/orders", middlewares.JWTMiddleware())
	{
		order.POST("", orderHandler.CreateOrder)
		order.GET("", orderHandler.GetAllOrders)
		order.GET("/:order_id", orderHandler.GetOrderByID)
		order.PUT("/:order_id", orderHandler.UpdateOrderByID)
		order.DELETE("/:order_id", orderHandler.DeleteOrderByID)

	}
}
