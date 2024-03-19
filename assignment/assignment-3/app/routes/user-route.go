package routes

import (
	"assignment-2/middlewares"
	"assignment-2/modules/user/handler"
	"assignment-2/modules/user/repository"
	"assignment-2/modules/user/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	account := r.Group("/accounts")
	{
		account.POST("register", userHandler.Register)
		account.POST("login", userHandler.Login)
	}

	user := r.Group("/users", middlewares.JWTMiddleware())
	{
		user.GET("/:user_id", userHandler.GetUserByID)
		user.PUT("/:user_id", userHandler.UpdateUserByID)
		user.DELETE("/:user_id", userHandler.DeleteUserByID)

	}
}
