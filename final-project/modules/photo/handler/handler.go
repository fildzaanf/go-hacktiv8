package handler

import (
	"final-project/middlewares"
	"final-project/modules/photo/dto/request"
	"final-project/modules/photo/dto/response"
	"final-project/modules/photo/entity"
	"final-project/utils/responses"
	"final-project/utils/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

type photoHandler struct {
	photoService entity.PhotoServiceInterface
}

func NewPhotoHandler(photoService entity.PhotoServiceInterface) *photoHandler {
	return &photoHandler{
		photoService: photoService,
	}
}


func (ph *photoHandler) CreatePhoto(c *gin.Context) {

	userID, role, errExtract := middlewares.VerifyToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}

	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}
	
	photoRequest := request.PhotoRequest{}

	errBind := c.ShouldBind(&photoRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	photoCore := request.PhotoRequestToPhotoCore(photoRequest)

	photoCore.UserID = userID

	photoData, errCreate := ph.photoService.CreatePhoto(photoCore)
	if errCreate != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errCreate.Error()))
		return
	}

	photoResponse := response.PhotoCoreToPhotoResponse(photoData)

	c.JSON(http.StatusCreated, responses.SuccessResponse("photo data created successfully", photoResponse))
}


func (ph *photoHandler) GetAllPhotos(c *gin.Context) {

	userID, role, errExtract := middlewares.VerifyToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}

	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}
	

	photosData, errGetAll := ph.photoService.GetAllPhotos(userID)
	if errGetAll != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse(errGetAll.Error()))
		return
	}


	if len(photosData) == 0 {
		c.JSON(http.StatusOK, responses.SuccessResponse("photos data empty", nil))
		return
	}

	photoResponse := response.ListPhotoCoreToPhotoResponse(photosData)

	c.JSON(http.StatusOK, responses.SuccessResponse("photos data retrieved successfully", photoResponse))
}

func (ph *photoHandler) GetPhotoByID(c *gin.Context) {
	photoID := c.Param("photo_id")

	userID, role, errExtract := middlewares.VerifyToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}

	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	photo, errGetID := ph.photoService.GetPhotoByID(photoID)
	if errGetID != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errGetID.Error()))
		return
	}

	if userID != photo.ID {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	photoResponse := response.PhotoCoreToPhotoResponse(photo)

	c.JSON(http.StatusOK, responses.SuccessResponse("photo data retrieved successfully", photoResponse))
}

func (ph *photoHandler) UpdatePhotoByID(c *gin.Context) {
	photoID := c.Param("photo_id")

	userID, role, errExtract := middlewares.VerifyToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}

	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	photo, errGetID := ph.photoService.GetPhotoByID(photoID)
	if errGetID != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errGetID.Error()))
		return
	}

	if userID != photo.ID {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	photoRequest := request.PhotoRequest{}

	errBind := c.ShouldBind(&photoRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	photoCore := request.PhotoRequestToPhotoCore(photoRequest)

	photo, errUpdate := ph.photoService.UpdatePhotoByID(photoID, photoCore)
	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errUpdate.Error()))
		return
	}

	photoResponse := response.PhotoCoreToPhotoResponse(photo)

	c.JSON(http.StatusOK, responses.SuccessResponse("photo data updated successfully", photoResponse))
}

func (ph *photoHandler) DeletePhotoByID(c *gin.Context) {
	photoID := c.Param("photo_id")

	userID, role, errExtract := middlewares.VerifyToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}
	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	photo, errGetID := ph.photoService.GetPhotoByID(photoID)
	if errGetID != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errGetID.Error()))
		return
	}

	if userID != photo.ID {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	errDelete := ph.photoService.DeletePhotoByID(photoID)
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse(errDelete.Error()))
		return
	}
	c.JSON(http.StatusOK, responses.SuccessResponse("photo data deleted successfully", nil))
}
