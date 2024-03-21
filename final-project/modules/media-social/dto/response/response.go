package response

import (
	"time"

	"gorm.io/gorm"
)

type MediaSocialResponse struct {
	ID             string
	Name           string
	MediaSocialURL string
	UserID         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}
