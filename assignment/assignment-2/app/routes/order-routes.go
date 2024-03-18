package routes

import (
	"assignment-2/handler"
	"assignment-2/repository"
	"assignment-2/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func OrderRoutes(r *gin.RouterGroup, db *gorm.DB) {
	orderRepository := repository.NewRepository(db)
	orderService := service.NewService(orderRepository)
	orderHandler := handler.NewHandler(orderService)

	order := r.Group("/orders")
	{
		order.POST("", orderHandler.CreateOrder)
		order.GET("", orderHandler.GetAllOrders)
		order.GET("/:order_id", orderHandler.GetOrderByID)
		order.PUT("/:order_id", orderHandler.UpdateOrderByID)
		order.DELETE("/:order_id", orderHandler.DeleteOrderByID)

	}
}
