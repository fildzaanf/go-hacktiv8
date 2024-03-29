package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	route := r.Group("")
	OrderRoutes(route, db)
	ItemRoutes(route, db)
}

