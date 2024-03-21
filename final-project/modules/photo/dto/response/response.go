package response

import (
	rc "final-project/modules/comment/dto/response"
	"time"

	"gorm.io/gorm"
)

type PhotoResponse struct {
	ID        string               `json:"id"`
	PhotoURL  string               `json:"photo_url"`
	Title     string               `json:"title"`
	Caption   string               `json:"caption"`
	UserID    string               `json:"user_id"`
	Comments  []rc.CommentResponse `json:"comments"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	DeletedAt gorm.DeletedAt       `json:"deleted_at"`
}
