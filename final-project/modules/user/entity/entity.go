package entity

import (
	ec "final-project/modules/comment/entity"
	em "final-project/modules/media-social/entity"
	ep "final-project/modules/photo/entity"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           string
	Username     string
	Fullname     string
	Email        string
	Password     string
	Age          int
	Role         string
	Photos       []ep.Photo
	Comments     []ec.Comment
	MediaSocials []em.MediaSocial
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}


