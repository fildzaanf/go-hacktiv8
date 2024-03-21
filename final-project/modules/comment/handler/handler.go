package handler

import (
	"final-project/middlewares"
	"final-project/modules/comment/dto/request"
	"final-project/modules/comment/dto/response"
	"final-project/modules/comment/entity"
	"final-project/utils/responses"
	"final-project/utils/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

type commentHandler struct {
	commentService entity.CommentServiceInterface
}

func NewCommentHandler(commentService entity.CommentServiceInterface) *commentHandler {
	return &commentHandler{
		commentService: commentService,
	}
}


func (ch *commentHandler) CreateComment(c *gin.Context) {

	userID, role, errExtract := middlewares.VerifyToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}

	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	CommentRequest := request.CommentRequest{}

	errBind := c.ShouldBind(&CommentRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	commentCore := request.CommentRequestToCommentCore(CommentRequest)

	commentCore.UserID = userID
	
	commentData, errCreate := ch.commentService.CreateComment(commentCore)
	if errCreate != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errCreate.Error()))
		return
	}

	commentResponse := response.CommentCoreToCommentResponse(commentData)

	c.JSON(http.StatusCreated, responses.SuccessResponse("comment data created successfully", commentResponse))
}



func (ch *commentHandler) GetAllComments(c *gin.Context) {

	userID, role, errExtract := middlewares.VerifyToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}
	
	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	commentsData, errGetAll := ch.commentService.GetAllComments(userID)
	if errGetAll != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse(errGetAll.Error()))
		return
	}

	if len(commentsData) == 0 {
		c.JSON(http.StatusOK, responses.SuccessResponse("comments data empty", nil))
		return
	}

	commentResponse := response.ListCommentCoreToCommentResponse(commentsData)

	c.JSON(http.StatusOK, responses.SuccessResponse("comments data retrieved successfully", commentResponse))
}


func (ch *commentHandler) GetCommentByID(c *gin.Context) {
	commentID := c.Param("comment_id")

	userID, role, errExtract := middlewares.VerifyToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}

	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	comment, errGetID := ch.commentService.GetCommentByID(commentID)
	if errGetID != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errGetID.Error()))
		return
	}

	if userID != comment.ID {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorizeds to access this resource"))
		return
	}

	commentResponse := response.CommentCoreToCommentResponse(comment)

	c.JSON(http.StatusOK, responses.SuccessResponse("comment data retrieved successfully", commentResponse))
}


func (ch *commentHandler) UpdateCommentByID(c *gin.Context) {
	commentID := c.Param("comment_id")

	userID, role, errExtract := middlewares.VerifyToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}

	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	comment, errGetID := ch.commentService.GetCommentByID(commentID)
	if errGetID != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errGetID.Error()))
		return
	}

	if userID != comment.ID {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorizeds to access this resource"))
		return
	}

	commentRequest := request.CommentRequest{}

	errBind := c.ShouldBind(&commentRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	commentCore := request.CommentRequestToCommentCore(commentRequest)

	comment, errUpdate := ch.commentService.UpdateCommentByID(commentID, commentCore)
	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errUpdate.Error()))
		return
	}

	commentResponse := response.CommentCoreToCommentResponse(comment)

	c.JSON(http.StatusOK, responses.SuccessResponse("comment data updated successfully", commentResponse))
}


func (ch *commentHandler) DeleteCommentByID(c *gin.Context) {
	commentID := c.Param("comment_id")

	userID, role, errExtract := middlewares.VerifyToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}
	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	comment, errGetID := ch.commentService.GetCommentByID(commentID)
	if errGetID != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errGetID.Error()))
		return
	}

	if userID != comment.ID {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorizeds to access this resource"))
		return
	}

	errDelete := ch.commentService.DeleteCommentByID(commentID)
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse(errDelete.Error()))
		return
	}
	c.JSON(http.StatusOK, responses.SuccessResponse("comment data deleted successfully", nil))
}
