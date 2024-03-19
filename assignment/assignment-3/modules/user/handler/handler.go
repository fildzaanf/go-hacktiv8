package handler

import (
	"assignment-2/modules/user/entity"
	"assignment-2/modules/user/dto/request"
	"assignment-2/modules/user/dto/response"
	"assignment-2/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService entity.UserServiceInterface
}

func NewUserHandler(us entity.UserServiceInterface) *userHandler {
	return &userHandler{
		userService: us,
	}
}

func (uh *userHandler) Register(c *gin.Context) {
	registerRequest := request.UserRegisterRequest{}

	errBind := c.ShouldBind(&registerRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	request := request.UserRegisterRequestToUserCore(registerRequest)

	_, errCreate := uh.userService.Register(request)
	if errCreate != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errCreate.Error()))
		return
	}

	response := response.UserCoreToUserRegisterResponse(request)
	c.JSON(http.StatusCreated, responses.SuccessResponse("register successfully", response))
}
