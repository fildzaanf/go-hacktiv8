package routes

import (
	"final-project/middlewares"
	"final-project/modules/comment/handler"
	"final-project/modules/comment/repository"
	"final-project/modules/comment/service"
	pr "final-project/modules/photo/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CommentRoutes(r *gin.RouterGroup, db *gorm.DB) {
	photoRepository := pr.NewPhotoRepository(db)
	commentRepository := repository.NewCommentRepository(db, photoRepository)
	commentService := service.NewCommentService(commentRepository)
	commentHandler := handler.NewCommentHandler(commentService)

	comment := r.Group("/comments", middlewares.JWTMiddleware(false))
	{
		comment.GET("", commentHandler.GetAllComments)
		comment.POST("/:photo_id", commentHandler.CreateComment)
		comment.GET("/:comment_id", commentHandler.GetCommentByID)
		comment.PUT("/:comment_id", commentHandler.UpdateCommentByID)
		comment.DELETE("/:comment_id", commentHandler.DeleteCommentByID)
	}
}
