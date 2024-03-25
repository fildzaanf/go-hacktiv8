package handler

import (
	"final-project/middlewares"
	"final-project/modules/media-social/dto/request"
	"final-project/modules/media-social/dto/response"
	"final-project/modules/media-social/entity"
	"final-project/utils/responses"
	"final-project/utils/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

type mediaSocialHandler struct {
	mediaSocialService entity.MediaSocialServiceInterface
}

func NewMediaSocialHandler(mediaSocialService entity.MediaSocialServiceInterface) *mediaSocialHandler {
	return &mediaSocialHandler{
		mediaSocialService: mediaSocialService,
	}
}

// CreateMediaSocial godoc
// @Summary Create a new media social entry
// @Description Create a new media social entry with the provided data
// @Tags media-social
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT access token"
// @Param mediaSocialRequest body request.MediaSocialRequest true "Media Social data"
// @Success 201 {object} response.MediaSocialResponse
// @Failure 400 {object} responses.TErrorResponse
// @Failure 401 {object} responses.TErrorResponse
// @Router /media-social [post]
func (mh *mediaSocialHandler) CreateMediaSocial(c *gin.Context) {
	userID, role, errExtract := middlewares.ExtractToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}

	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	MediaSocialRequest := request.MediaSocialRequest{}

	errBind := c.ShouldBind(&MediaSocialRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	mediaSocialCore := request.MediaSocialRequestToMediaSocialCore(MediaSocialRequest)
	
	mediaSocialCore.UserID = userID

	mediaSocialData, errCreate := mh.mediaSocialService.CreateMediaSocial(mediaSocialCore)
	if errCreate != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errCreate.Error()))
		return
	}

	mediaSocialResponse := response.MediaSocialCoreToMediaSocialResponse(mediaSocialData)

	c.JSON(http.StatusCreated, responses.SuccessResponse("media social data created successfully", mediaSocialResponse))
}

// GetAllMediaSocials godoc
// @Summary Get all media social entries
// @Description Retrieve all media social entries
// @Tags media-social
// @Produce json
// @Param Authorization header string true "JWT access token"
// @Success 200 {object} response.MediaSocialResponse
// @Failure 401 {object} responses.TErrorResponse
// @Router /media-social [get]
func (mh *mediaSocialHandler) GetAllMediaSocials(c *gin.Context) {

	userID, role, errExtract := middlewares.ExtractToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}
	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	mediaSocialsData, errGetAll := mh.mediaSocialService.GetAllMediaSocials(userID)
	if errGetAll != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse(errGetAll.Error()))
		return
	}

	if len(mediaSocialsData) == 0 {
		c.JSON(http.StatusOK, responses.SuccessResponse("media socials data empty", nil))
		return
	}

	mediaSocialResponse := response.ListMediaSocialCoreToMediaSocialResponse(mediaSocialsData)

	c.JSON(http.StatusOK, responses.SuccessResponse("media socials data retrieved successfully", mediaSocialResponse))
}

// GetMediaSocialByID godoc
// @Summary Get media social entry by ID
// @Description Retrieve media social entry details by media social ID
// @Tags media-social
// @Produce json
// @Param Authorization header string true "JWT access token"
// @Param medsos_id path string true "Media Social ID"
// @Success 200 {object} response.MediaSocialResponse
// @Failure 400 {object} responses.TErrorResponse
// @Failure 401 {object} responses.TErrorResponse
// @Router /media-social/{medsos_id} [get]
func (mh *mediaSocialHandler) GetMediaSocialByID(c *gin.Context) {
	mediaSocialID := c.Param("medsos_id")

	_, role, errExtract := middlewares.ExtractToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}

	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	mediaSocial, errGetID := mh.mediaSocialService.GetMediaSocialByID(mediaSocialID)
	if errGetID != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errGetID.Error()))
		return
	}

	mediaSocialResponse := response.MediaSocialCoreToMediaSocialResponse(mediaSocial)

	c.JSON(http.StatusOK, responses.SuccessResponse("media social data retrieved successfully", mediaSocialResponse))
}

// UpdateMediaSocialByID godoc
// @Summary Update media social entry by ID
// @Description Update media social entry details by media social ID
// @Tags media-social
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT access token"
// @Param medsos_id path string true "Media Social ID"
// @Param mediaSocialRequest body request.MediaSocialRequest true "Media Social details"
// @Success 200 {object} response.MediaSocialResponse
// @Failure 400 {object} responses.TErrorResponse
// @Failure 401 {object} responses.TErrorResponse
// @Router /media-social/{medsos_id} [put]
func (mh *mediaSocialHandler) UpdateMediaSocialByID(c *gin.Context) {
	mediaSocialID := c.Param("medsos_id")

	userID, role, errExtract := middlewares.ExtractToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}

	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	mediaSocial, errGetID := mh.mediaSocialService.GetMediaSocialByID(mediaSocialID)
	if errGetID != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errGetID.Error()))
		return
	}

	if userID != mediaSocial.UserID {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	mediaSocialRequest := request.MediaSocialRequest{}

	errBind := c.ShouldBind(&mediaSocialRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	mediaSocialCore := request.MediaSocialRequestToMediaSocialCore(mediaSocialRequest)

	mediaSocial, errUpdate := mh.mediaSocialService.UpdateMediaSocialByID(mediaSocialID, mediaSocialCore)
	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errUpdate.Error()))
		return
	}

	mediaSocialResponse := response.MediaSocialCoreToMediaSocialResponse(mediaSocial)

	c.JSON(http.StatusOK, responses.SuccessResponse("media social data updated successfully", mediaSocialResponse))
}

// DeleteMediaSocialByID godoc
// @Summary Delete media social entry by ID
// @Description Delete media social entry by media social ID
// @Tags media-social
// @Produce json
// @Param Authorization header string true "JWT access token"
// @Param medsos_id path string true "Media Social ID"
// @Success 200 {object} responses.TSuccessResponse
// @Failure 401 {object} responses.TErrorResponse
// @Router /media-social/{medsos_id} [delete]
func (mh *mediaSocialHandler) DeleteMediaSocialByID(c *gin.Context) {
	mediaSocialID := c.Param("medsos_id")

	userID, role, errExtract := middlewares.ExtractToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}
	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	mediaSocial, errGetID := mh.mediaSocialService.GetMediaSocialByID(mediaSocialID)
	if errGetID != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errGetID.Error()))
		return
	}

	if userID != mediaSocial.UserID {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	errDelete := mh.mediaSocialService.DeleteMediaSocialByID(mediaSocialID)
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse(errDelete.Error()))
		return
	}
	c.JSON(http.StatusOK, responses.SuccessResponse("media social data deleted successfully", nil))
}
