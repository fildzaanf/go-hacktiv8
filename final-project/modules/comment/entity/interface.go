package entity

import "github.com/gin-gonic/gin"

type CommentRepositoryInterface interface {
	CreateComment(photoID string, commentCore Comment) (Comment, error)
	GetCommentByID(commentID string) (Comment, error)
	GetAllComments(userID string) ([]Comment, error)
	UpdateCommentByID(commentID string, commentCore Comment) (Comment, error)
	DeleteCommentByID(commentID string) error
}

type CommentServiceInterface interface {
	CreateComment(photoID string, commentCore Comment) (Comment, error)
	GetCommentByID(commentID string) (Comment, error)
	GetAllComments(userID string) ([]Comment, error)
	UpdateCommentByID(commentID string, commentCore Comment) (Comment, error)
	DeleteCommentByID(commentID string) error
}

type CommentHandlerInterface interface {
	CreateComment(c gin.Context)
	GetAllComments(c gin.Context)
	GetCommentByID(c gin.Context)
	UpdateCommentByID(c gin.Context)
	DeleteCommentByID(c gin.Context)
}
