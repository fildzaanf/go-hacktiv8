package handler

import (
	"final-project/middlewares"
	"final-project/modules/user/dto/request"
	"final-project/modules/user/dto/response"
	"final-project/modules/user/entity"
	"final-project/utils/constant"
	"final-project/utils/responses"
	"net/http"
	"strings"

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
// @Param userRequest body request.UserRegisterRequest true "User data for registration"
// @Success 201 {object} response.UserRegisterResponse
// @Failure 400 {object} responses.TErrorResponse
// @Router /accounts/register [post]
func (uh *userHandler) Register(c *gin.Context) {
	userRequest := request.UserRegisterRequest{}

	errBind := c.ShouldBind(&userRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	userCore := request.UserRegisterRequestToUserCore(userRequest)

	user, errCreate := uh.userService.Register(userCore)
	if errCreate != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errCreate.Error()))
		return
	}

	userResponse := response.UserCoreToUserRegisterResponse(user)
	c.JSON(http.StatusCreated, responses.SuccessResponse("register successfully", userResponse))
}

// Login godoc
// @Summary Login user
// @Description Login user with provided email and password
// @Tags accounts
// @Accept json
// @Produce json
// @Param userRequest body request.UserLoginRequest true "User login credentials"
// @Success 200 {object} response.UserLoginResponse
// @Failure 400 {object} responses.TErrorResponse
// @Router /accounts/login [post]
func (uh *userHandler) Login(c *gin.Context) {
	userRequest := request.UserLoginRequest{}

	errBind := c.ShouldBind(&userRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	user, token, errLogin := uh.userService.Login(userRequest.Email, userRequest.Password)
	if errLogin != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errLogin.Error()))
		return
	}

	userResponse := response.UserCoreToUserLoginResponse(user, token)
	c.JSON(http.StatusOK, responses.SuccessResponse("login successfully", userResponse))
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

	_, role, errExtract := middlewares.ExtractToken(c)
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

	userResponse := response.ListUserCoreToUserResponse(users)

	c.JSON(http.StatusOK, responses.SuccessResponse("data retrieved successfully", userResponse))
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
	userID := c.Param("user_id")

	_, role, errExtract := middlewares.ExtractToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}
	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	// if tokenUserID != pathUserID {
	// 	c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorizeds to access this resource"))
	// 	return
	// }

	user, errGetID := uh.userService.GetUserByID(userID)
	if errGetID != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errGetID.Error()))
		return
	}

	userResponse := response.UserCoreToUserResponse(user)
	c.JSON(http.StatusOK, responses.SuccessResponse("data retrieved successfully", userResponse))
}

// UpdateUserByID godoc
// @Summary Update user by ID
// @Description Update user details by user ID
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT access token"
// @Param user_id path string true "User ID"
// @Param userRequest body request.UserUpdateProfileRequest true "User details"
// @Success 200 {object} response.UserResponse
// @Failure 400 {object} responses.TErrorResponse
// @Failure 401 {object} responses.TErrorResponse
// @Router /users/{user_id} [put]
func (uh *userHandler) UpdateUserByID(c *gin.Context) {
	pathUserID := c.Param("user_id")

	tokenUserID, role, errExtract := middlewares.ExtractToken(c)
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

	userRequest := request.UserUpdateProfileRequest{}

	errBind := c.ShouldBind(&userRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	userCore := request.UserUpdateProfileRequestToUserCore(userRequest)

	user, errUpdate := uh.userService.UpdateUserByID(pathUserID, userCore)
	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errUpdate.Error()))
		return
	}

	userResponse := response.UserCoreToUserResponse(user)
	c.JSON(http.StatusOK, responses.SuccessResponse("data updated successfully", userResponse))
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

	tokenUserID, role, errExtract := middlewares.ExtractToken(c)
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

// ForgotPassword godoc
// @Summary Forgot password
// @Description Initiate password reset process by sending OTP to user's email
// @Tags passwords
// @Accept json
// @Produce json
// @Param userRequest body request.UserSendOTPRequest true "User email for sending OTP"
// @Success 200 {object} responses.TSuccessResponse
// @Failure 400 {object} responses.TErrorResponse
// @Router /password/forgot-password [post]
func (uh *userHandler) ForgotPassword(c *gin.Context) {
	userRequest := request.UserSendOTPRequest{}

	errBind := c.Bind(&userRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	userCore := request.UserSendOTPRequestToUserCore(userRequest)

	errSendOTP := uh.userService.SendOTP(userCore.Email)
	if errSendOTP != nil {
		if strings.Contains(errSendOTP.Error(), "email not found") {
			c.JSON(http.StatusNotFound, responses.ErrorResponse(errSendOTP.Error()))
			return
		}
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errSendOTP.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse("otp sent successfully", nil))
}

// VerifyOTP godoc
// @Summary Verify OTP
// @Description Verify OTP sent to user's email for password reset
// @Tags passwords
// @Accept json
// @Produce json
// @Param userRequest body request.UserVerifyOTPRequest true "User email and OTP for verification"
// @Success 200 {string} string "OTP verification successfully"
// @Success 200 {object} response.UserResponse
// @Failure 400 {object} responses.TErrorResponse
// @Router /password/verify-otp [post]
func (uh *userHandler) VerifyOTP(c *gin.Context) {
	userRequest := request.UserVerifyOTPRequest{}

	errBind := c.Bind(&userRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	userCore := request.UserVerifyOTPRequestToUserCore(userRequest)

	token, errVerify := uh.userService.VerifyOTP(userCore.Email, userCore.OTP)
	if errVerify != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse("otp verification failed "+errVerify.Error()))
		return
	}

	userResponse :=  response.UserCoreToUserVerifyOTPResponse(token)

	c.JSON(http.StatusOK, responses.SuccessResponse("otp verification successfully", userResponse))
}

// NewPassword godoc
// @Summary Set new password
// @Description Set new password after verifying OTP
// @Tags passwords
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT access token"
// @Param userRequest body request.UserNewPasswordRequest true "New password details"
// @Success 200 {object} responses.TSuccessResponse
// @Failure 400 {object} responses.TErrorResponse
// @Failure 401 {object} responses.TErrorResponse
// @Router /password/change-password [post]
func (uh *userHandler) NewPassword(c *gin.Context) {
	userRequest := request.UserNewPasswordRequest{}

	errBind := c.Bind(&userRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	email, errExtract := middlewares.ExtractVerifyToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}

	userCore := request.UserNewPasswordRequestToUserCore(userRequest)

	errUpdate := uh.userService.NewPassword(email, userCore)
	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errUpdate.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse("password updated successfully", nil))
}
