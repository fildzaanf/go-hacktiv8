package routes

import (
	"final-project/middlewares"
	"final-project/modules/photo/handler"
	"final-project/modules/photo/repository"
	"final-project/modules/photo/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PhotoRoutes(r *gin.RouterGroup, db *gorm.DB) {
	photoRepository := repository.NewPhotoRepository(db)
	photoService := service.NewPhotoService(photoRepository)
	photoHandler := handler.NewPhotoHandler(photoService)

	photo := r.Group("/photos", middlewares.JWTMiddleware(false))
	{
		photo.POST("", photoHandler.CreatePhoto)
		photo.GET("", photoHandler.GetAllPhotos)
		photo.GET("/:photo_id", photoHandler.GetPhotoByID)
		photo.PUT("/:photo_id", photoHandler.UpdatePhotoByID)
		photo.DELETE("/:photo_id", photoHandler.DeletePhotoByID)

	}
}
