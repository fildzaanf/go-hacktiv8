package entity

import "github.com/gin-gonic/gin"

type MediaSocialRepositoryInterface interface {
	CreateMediaSocial(mediaSocialCore MediaSocial) (MediaSocial, error)
	GetMediaSocialByID(mediaSocialID string) (MediaSocial, error)
	GetAllMediaSocials(userID string) ([]MediaSocial, error)
	UpdateMediaSocialByID(mediaSocialID string, mediaSocialCore MediaSocial) (MediaSocial, error)
	DeleteMediaSocialByID(mediaSocialID string) error
}

type MediaSocialServiceInterface interface {
	CreateMediaSocial(mediaSocialCore MediaSocial) (MediaSocial, error)
	GetMediaSocialByID(mediaSocialID string) (MediaSocial, error)
	GetAllMediaSocials(userID string) ([]MediaSocial, error)
	UpdateMediaSocialByID(mediaSocialID string, mediaSocialCore MediaSocial) (MediaSocial, error)
	DeleteMediaSocialByID(mediaSocialID string) error
}

type MediaSocialHandlerInterface interface {
	CreateMediaSocial(c gin.Context)
	GetAllMediaSocials(c gin.Context)
	GetMediaSocialByID(c gin.Context)
	UpdateMediaSocialByID(c gin.Context)
	DeleteMediaSocialByID(c gin.Context)
}
