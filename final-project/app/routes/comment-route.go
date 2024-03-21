package routes

import (
	"final-project/middlewares"
	"final-project/modules/comment/handler"
	"final-project/modules/comment/repository"
	"final-project/modules/comment/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CommentRoutes(r *gin.RouterGroup, db *gorm.DB) {
	commentRepository := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepository)
	commentHandler := handler.NewCommentHandler(commentService)

	comment := r.Group("/comments", middlewares.JWTMiddleware())
	{
		comment.POST("", commentHandler.CreateComment)
		comment.GET("", commentHandler.GetAllComments)
		comment.GET(":comment_id", commentHandler.GetCommentByID)
		comment.PUT(":comment_id", commentHandler.UpdateCommentByID)
		comment.DELETE(":comment_id", commentHandler.DeleteCommentByID)

	}
}
