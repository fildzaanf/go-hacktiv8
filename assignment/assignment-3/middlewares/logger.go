package middlewares

import "github.com/gin-gonic/gin"

func Logger(router *gin.Engine) {
	router.Use(gin.Logger())
}
