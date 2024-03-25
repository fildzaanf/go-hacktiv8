package handler

import (
	"final-project/middlewares"
	"final-project/modules/comment/dto/request"
	"final-project/modules/comment/dto/response"
	"final-project/modules/comment/entity"
	"final-project/utils/constant"
	"final-project/utils/responses"
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

// CreateComment godoc
// @Summary Create a new comment
// @Description Create a new comment with the provided comment data
// @Tags comments
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT access token"
// @Param photo_id path string true "Photo ID"
// @Param commentRequest body request.CommentRequest true "Comment data"
// @Success 201 {object} response.CommentResponse
// @Success 201 {object} responses.TSuccessResponse
// @Failure 400 {object} responses.TErrorResponse
// @Failure 401 {object} responses.TErrorResponse
// @Router /comments/{photo_id} [post]
func (ch *commentHandler) CreateComment(c *gin.Context) {
	photoID := c.Param("photo_id")

	userID, role, errExtract := middlewares.ExtractToken(c)
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

	commentData, errCreate := ch.commentService.CreateComment(photoID, commentCore)
	if errCreate != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errCreate.Error()))
		return
	}

	commentResponse := response.CommentCoreToCommentResponse(commentData)

	c.JSON(http.StatusCreated, responses.SuccessResponse("comment data created successfully", commentResponse))
}

// GetAllComments godoc
// @Summary Get all comments
// @Description Retrieve all comments
// @Tags comments
// @Produce json
// @Param Authorization header string true "JWT access token"
// @Success 200 {object} response.CommentResponse
// @Success 200 {object} responses.TSuccessResponse
// @Failure 401 {object} responses.TErrorResponse
// @Failure 500 {object} responses.TErrorResponse
// @Router /comments [get]
func (ch *commentHandler) GetAllComments(c *gin.Context) {

	userID, role, errExtract := middlewares.ExtractToken(c)
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

// GetCommentByID godoc
// @Summary Get comment by ID
// @Description Retrieve comment details by comment ID
// @Tags comments
// @Produce json
// @Param Authorization header string true "JWT access token"
// @Param comment_id path string true "Comment ID"
// @Success 200 {object} response.CommentResponse
// @Success 200 {object} responses.TSuccessResponse
// @Failure 400 {object} responses.TErrorResponse
// @Failure 401 {object} responses.TErrorResponse
// @Router /comments/{comment_id} [get]
func (ch *commentHandler) GetCommentByID(c *gin.Context) {
	commentID := c.Param("comment_id")

	_, role, errExtract := middlewares.ExtractToken(c)
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

	commentResponse := response.CommentCoreToCommentResponse(comment)

	c.JSON(http.StatusOK, responses.SuccessResponse("comment data retrieved successfully", commentResponse))
}

// UpdateCommentByID godoc
// @Summary Update comment by ID
// @Description Update comment details by comment ID
// @Tags comments
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT access token"
// @Param comment_id path string true "Comment ID"
// @Param commentRequest body request.CommentRequest true "Comment details"
// @Success 200 {object} response.CommentResponse
// @Success 200 {object} responses.TSuccessResponse
// @Failure 400 {object} responses.TErrorResponse
// @Failure 401 {object} responses.TErrorResponse
// @Router /comments/{comment_id} [put]
func (ch *commentHandler) UpdateCommentByID(c *gin.Context) {
	commentID := c.Param("comment_id")

	userID, role, errExtract := middlewares.ExtractToken(c)
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

	if userID != comment.UserID {
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

// DeleteCommentByID godoc
// @Summary Delete comment by ID
// @Description Delete comment by comment ID
// @Tags comments
// @Produce json
// @Param Authorization header string true "JWT access token"
// @Param comment_id path string true "Comment ID"
// @Success 200 {object} responses.TSuccessResponse
// @Failure 401 {object} responses.TErrorResponse
// @Failure 400 {object} responses.TErrorResponse
// @Failure 500 {object} responses.TErrorResponse
// @Router /comments/{comment_id} [delete]
func (ch *commentHandler) DeleteCommentByID(c *gin.Context) {
	commentID := c.Param("comment_id")

	userID, role, errExtract := middlewares.ExtractToken(c)
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

	if userID != comment.UserID {
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
