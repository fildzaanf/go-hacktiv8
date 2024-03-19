package routes

import (
	"assignment-2/modules/order/handler"
	"assignment-2/modules/order/repository"
	"assignment-2/modules/order/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ItemRoutes(r *gin.RouterGroup, db *gorm.DB) {
	itemRepository := repository.NewOrderRepository(db)
	itemService := service.NewOrderService(itemRepository)
	itemHandler := handler.NewOrderHandler(itemService)

	item := r.Group("/items")
	{
		item.POST("", itemHandler.CreateItem)
		item.GET("", itemHandler.GetAllItems)
		item.GET("/:item_id", itemHandler.GetItemByID)
		item.PUT("/:item_id", itemHandler.UpdateItemByID)
		item.DELETE("/:item_id", itemHandler.DeleteItemByID)
	}
}
