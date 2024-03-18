package routes

import (
	"assignment-2/handler"
	"assignment-2/repository"
	"assignment-2/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ItemRoutes(r *gin.RouterGroup, db *gorm.DB) {
	itemRepository := repository.NewRepository(db)
	itemService := service.NewService(itemRepository)
	itemHandler := handler.NewHandler(itemService)

	item := r.Group("/items")
	{
		item.POST("", itemHandler.CreateItem)
		item.GET("", itemHandler.GetAllItems)
		item.GET("/:item_id", itemHandler.GetItemByID)
		item.PUT("/:item_id", itemHandler.UpdateItemByID)
		item.DELETE("/:item_id", itemHandler.DeleteItemByID)
	}
}
