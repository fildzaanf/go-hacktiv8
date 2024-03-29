package entity

import "github.com/gin-gonic/gin"

type UserRepositoryInterface interface {
	Register(userCore User) (User, error)
	Login(email, password string) (User, error)
	GetUserByID(userID string) (User, error)
	GetUserByEmail(email string) (User, error)
	UpdateUserByID(userID string, userCore User) (User,error)
	DeleteUserByID(userID string) error
}

type UserServiceInterface interface {
	Register(userCore User) (User, error)
	Login(email, password string) (User, string, error)
	GetUserByEmail(userID string) (User, error)
	GetUserByID(userID string) (User, error)

	UpdateUserByID(userID string, userCore User) (User,error)
	DeleteUserByID(userID string) error
}

type UserHandlerInterface interface {
	Register(c gin.Context)
	Login(c gin.Context)
	GetUserByID(c gin.Context)
	UpdateUserByID(c gin.Context)
	DeletedUserByID(c gin.Context)
}
