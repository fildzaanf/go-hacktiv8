package routes

import (
	"final-project/middlewares"
	"final-project/modules/media-social/handler"
	"final-project/modules/media-social/repository"
	"final-project/modules/media-social/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MediaSocialRoutes(r *gin.RouterGroup, db *gorm.DB) {
	medsosRepository := repository.NewMediaSocialRepository(db)
	medsosService := service.NewMediaSocialService(medsosRepository)
	medsosHandler := handler.NewMediaSocialHandler(medsosService)

	medsos := r.Group("/media-socials", middlewares.JWTMiddleware(false))
	{
		medsos.POST("", medsosHandler.CreateMediaSocial)
		medsos.GET("", medsosHandler.GetAllMediaSocials)
		medsos.GET("/:medsos_id", medsosHandler.GetMediaSocialByID)
		medsos.PUT("/:medsos_id", medsosHandler.UpdateMediaSocialByID)
		medsos.DELETE("/:medsos_id", medsosHandler.DeleteMediaSocialByID)
	}
}
