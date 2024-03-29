package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	route := r.Group("")
	UserRoutes(route, db)
	OrderRoutes(route, db)	
}

