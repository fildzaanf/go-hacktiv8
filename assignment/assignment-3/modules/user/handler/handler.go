package handler

import (
	"assignment-2/middlewares"
	"assignment-2/modules/user/dto/request"
	"assignment-2/modules/user/dto/response"
	"assignment-2/modules/user/entity"
	"assignment-2/utils/responses"
	"assignment-2/utils/constant"
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

func (uh *userHandler) Login(c *gin.Context) {
	loginRequest := request.UserLoginRequest{}

	errBind := c.ShouldBind(&loginRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	request, token, errLogin := uh.userService.Login(loginRequest.Email, loginRequest.Password)
	if errLogin != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errLogin.Error()))
		return
	}

	response := response.UserCoreToUserLoginResponse(request, token)
	c.JSON(http.StatusOK, responses.SuccessResponse("login successfully", response))
}

func (uh *userHandler) GetUserByID(c *gin.Context) {
	userID, role, errExtract := middlewares.VerifyToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}
	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	result, errGetID := uh.userService.GetUserByID(userID)
	if errGetID != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errGetID.Error()))
		return
	}

	response := response.UserCoreToUserResponse(result)
	c.JSON(http.StatusOK, responses.SuccessResponse("data retrieved successfully", response))
}

func (uh *userHandler) UpdateUserByID(c *gin.Context) {
	userRequest := request.UserRequest{}

	errBind := c.ShouldBind(&userRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	userID, role, errExtract := middlewares.VerifyToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}
	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	request := request.UserRequestToUserCore(userRequest)

	errUpdate := uh.userService.UpdateUserByID(userID, request)
	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errUpdate.Error()))
		return
	}

	response := response.UserCoreToUserResponse(request)
	c.JSON(http.StatusOK, responses.SuccessResponse("data updated successfully", response))
}