package model

import (
	"time"

	"gorm.io/gorm"
)

type MediaSocial struct {
	ID             string `gorm:"primarykey"`
	Name           string `gorm:"not null"`
	MediaSocialURL string `gorm:"not null"`
	UserID         string `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
