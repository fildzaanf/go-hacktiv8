package handler

import (
	"final-project/middlewares"
	"final-project/modules/user/dto/request"
	"final-project/modules/user/dto/response"
	"final-project/modules/user/entity"
	"final-project/utils/responses"
	"final-project/utils/constant"
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

// Register godoc
// @Summary Register a new user
// @Description Register a new user with the provided data
// @Tags accounts
// @Accept json
// @Produce json
// @Param userRegisterRequest body request.UserRegisterRequest true "User data for registration"
// @Success 201 {object} response.UserRegisterResponse
// @Failure 400 {object} responses.TErrorResponse
// @Router /accounts/register [post]
func (uh *userHandler) Register(c *gin.Context) {
	registerRequest := request.UserRegisterRequest{}

	errBind := c.ShouldBind(&registerRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	request := request.UserRegisterRequestToUserCore(registerRequest)

	user, errCreate := uh.userService.Register(request)
	if errCreate != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errCreate.Error()))
		return
	}

	response := response.UserCoreToUserRegisterResponse(user)
	c.JSON(http.StatusCreated, responses.SuccessResponse("register successfully", response))
}

// Login godoc
// @Summary Login user
// @Description Login user with provided email and password
// @Tags accounts
// @Accept json
// @Produce json
// @Param userLoginRequest body request.UserLoginRequest true "User login credentials"
// @Success 200 {object} response.UserLoginResponse
// @Failure 400 {object} responses.TErrorResponse
// @Router /accounts/login [post]
func (uh *userHandler) Login(c *gin.Context) {
	loginRequest := request.UserLoginRequest{}

	errBind := c.ShouldBind(&loginRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	user, token, errLogin := uh.userService.Login(loginRequest.Email, loginRequest.Password)
	if errLogin != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errLogin.Error()))
		return
	}

	response := response.UserCoreToUserLoginResponse(user, token)
	c.JSON(http.StatusOK, responses.SuccessResponse("login successfully", response))
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieve all users
// @Tags users
// @Produce json
// @Param Authorization header string true "JWT access token"
// @Success 200 {object} response.UserResponse
// @Failure 401 {object} responses.TErrorResponse
// @Router /users [get]
func (uh *userHandler) GetAllUsers(c *gin.Context) {

	_, role, errExtract := middlewares.VerifyToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}
	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	users, errGetAll := uh.userService.GetAllUsers()
	if errGetAll != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse(errGetAll.Error()))
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusOK, responses.SuccessResponse("data is empty", nil))
		return
	}

	response := response.ListUserCoreToUserResponse(users)

	c.JSON(http.StatusOK, responses.SuccessResponse("data retrieved successfully", response))
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Retrieve user details by user ID
// @Tags users
// @Produce json
// @Param Authorization header string true "JWT access token"
// @Param user_id path string true "User ID"
// @Success 200 {object} response.UserResponse
// @Failure 400 {object} responses.TErrorResponse
// @Failure 401 {object} responses.TErrorResponse
// @Router /users/{user_id} [get]
func (uh *userHandler) GetUserByID(c *gin.Context) {
	pathUserID := c.Param("user_id")

	tokenUserID, role, errExtract := middlewares.VerifyToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}
	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	if tokenUserID != pathUserID {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorizeds to access this resource"))
		return
	}

	user, errGetID := uh.userService.GetUserByID(pathUserID)
	if errGetID != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errGetID.Error()))
		return
	}

	response := response.UserCoreToUserResponse(user)
	c.JSON(http.StatusOK, responses.SuccessResponse("data retrieved successfully", response))
}

// UpdateUserByID godoc
// @Summary Update user by ID
// @Description Update user details by user ID
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT access token"
// @Param user_id path string true "User ID"
// @Param userRequest body request.UserRequest true "User details"
// @Success 200 {object} response.UserResponse
// @Failure 400 {object} responses.TErrorResponse
// @Failure 401 {object} responses.TErrorResponse
// @Router /users/{user_id} [put]
func (uh *userHandler) UpdateUserByID(c *gin.Context) {
	pathUserID := c.Param("user_id")

	tokenUserID, role, errExtract := middlewares.VerifyToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}
	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	if tokenUserID != pathUserID {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	userRequest := request.UserRequest{}

	errBind := c.ShouldBind(&userRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	request := request.UserRequestToUserCore(userRequest)

	user, errUpdate := uh.userService.UpdateUserByID(pathUserID, request)
	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errUpdate.Error()))
		return
	}

	response := response.UserCoreToUserResponse(user)
	c.JSON(http.StatusOK, responses.SuccessResponse("data updated successfully", response))
}

// DeleteUserByID godoc
// @Summary Delete user by ID
// @Description Delete user by user ID
// @Tags users
// @Produce json
// @Param Authorization header string true "JWT access token"
// @Param user_id path string true "User ID"
// @Success 200 {object} responses.TSuccessResponse
// @Failure 401 {object} responses.TErrorResponse
// @Router /users/{user_id} [delete]
func (uh *userHandler) DeleteUserByID(c *gin.Context) {
	pathUserID := c.Param("user_id")

	tokenUserID, role, errExtract := middlewares.VerifyToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}
	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}
	if tokenUserID != pathUserID {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	errDelete := uh.userService.DeleteUserByID(pathUserID)
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse(errDelete.Error()))
		return
	}
	c.JSON(http.StatusOK, responses.SuccessResponse("data deleted successfully", nil))
}
