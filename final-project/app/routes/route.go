package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	base := r.Group("")
	UserRoutes(base, db)
	PhotoRoutes(base, db)	
	CommentRoutes(base, db)
	MediaSocialRoutes(base, db)
}

