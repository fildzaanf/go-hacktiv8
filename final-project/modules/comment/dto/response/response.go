package response

import (
	"time"

	"gorm.io/gorm"
)

type CommentResponse struct {
	ID        string         `json:"id"`
	UserID    string         `json:"user_id"`
	PhotoID   string         `json:"photo_id"`
	Message   string         `json:"message"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
