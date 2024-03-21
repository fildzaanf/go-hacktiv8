package entity

import "github.com/gin-gonic/gin"

type PhotoRepositoryInterface interface {
	CreatePhoto(photoCore Photo) (Photo, error)
	GetPhotoByID(photoID string) (Photo, error)
	GetAllPhotos(userID string) ([]Photo, error)
	UpdatePhotoByID(photoID string, photoCore Photo) (Photo, error)
	DeletePhotoByID(photoID string) error
}

type PhotoServiceInterface interface {
	CreatePhoto(photoCore Photo) (Photo, error)
	GetPhotoByID(photoID string) (Photo, error)
	GetAllPhotos(userID string) ([]Photo, error)
	UpdatePhotoByID(photoID string, photoCore Photo) (Photo, error)
	DeletePhotoByID(photoID string) error
}

type PhotoHandlerInterface interface {
	CreatePhoto(c gin.Context)
	GetAllPhotos(c gin.Context)
	GetPhotoByID(c gin.Context)
	UpdatePhotoByID(c gin.Context)
	DeletePhotoByID(c gin.Context)
}
