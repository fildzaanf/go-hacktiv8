package entity

import "github.com/gin-gonic/gin"

type UserRepositoryInterface interface {
	Register(userCore User) (User, error)
	Login(email, password string) (User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(userID string) (User, error)
	GetUserByEmail(email string) (User, error)
	GetUserByUsername(username string) (User, error)
	UpdateUserByID(userID string, userCore User) (User, error)
	DeleteUserByID(userID string) error
	NewPassword(email string, userCore User) (User, error)
	SendOTP(email string, otp string, expired int64) (User, error)
	VerifyOTP(email, otp string) (User, error)
	ResetOTP(otp string) (User, error)
}

type UserServiceInterface interface {
	Register(userCore User) (User, error)
	Login(email, password string) (User, string, error)
	GetAllUsers() ([]User, error)
	GetUserByEmail(userID string) (User, error)
	GetUserByID(userID string) (User, error)
	GetUserByUsername(username string) (User, error)
	UpdateUserByID(userID string, userCore User) (User, error)
	DeleteUserByID(userID string) error
	NewPassword(email string, userCore User) error
	SendOTP(email string) error
	VerifyOTP(email, otp string) (string, error)
}

type UserHandlerInterface interface {
	Register(c gin.Context)
	Login(c gin.Context)
	GetAllUsers(c gin.Context)
	GetUserByID(c gin.Context)
	UpdateUserByID(c gin.Context)
	DeletedUserByID(c gin.Context)
	ForgotPassword(c gin.Context)
	NewPassword(c gin.Context)
	VerifyOTP(c gin.Context)
}
