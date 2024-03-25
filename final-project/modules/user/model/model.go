package model

import (
	cm "final-project/modules/comment/model"
	sm "final-project/modules/media-social/model"
	pm "final-project/modules/photo/model"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            string `gorm:"primarykey"`
	Username      string `gorm:"not null"`
	Fullname      string
	Email         string `gorm:"not null"`
	Password      string `gorm:"not null"`
	Age           int
	Role          string `gorm:"not null"`
	OTP           string `gorm:"not null"`
	OTPExpiration int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt   `gorm:"index"`
	Photos        []pm.Photo       `gorm:"foreignKey:UserID;references:ID"`
	Comments      []cm.Comment     `gorm:"foreignKey:UserID;references:ID"`
	MediaSocials  []sm.MediaSocial `gorm:"foreignKey:UserID;references:ID"`
}
