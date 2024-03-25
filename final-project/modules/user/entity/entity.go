package entity

import (
	ec "final-project/modules/comment/entity"
	em "final-project/modules/media-social/entity"
	ep "final-project/modules/photo/entity"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              string
	Username        string
	Fullname        string
	Email           string
	Password        string
	ConfirmPassword string
	NewPassword     string
	Age             int
	Role            string
	OTP             string
	OTPExpiration   int64
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
	Photos          []ep.Photo
	Comments        []ec.Comment
	MediaSocials    []em.MediaSocial
}
