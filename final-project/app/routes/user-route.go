package routes

import (
	"final-project/middlewares"
	"final-project/modules/user/handler"
	"final-project/modules/user/repository"
	"final-project/modules/user/service"

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
		user.GET("", userHandler.GetAllUsers)
		user.GET("/:user_id", userHandler.GetUserByID)
		user.PUT("/:user_id", userHandler.UpdateUserByID)
		user.DELETE("/:user_id", userHandler.DeleteUserByID)

	}
}
