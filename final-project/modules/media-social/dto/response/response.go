package response

import (
	"time"

	"gorm.io/gorm"
)

type MediaSocialResponse struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	MediaSocialURL string         `json:"media_social_url"`
	UserID         string         `json:"user_id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
}
