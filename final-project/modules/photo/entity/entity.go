package entity

import (
	ec "final-project/modules/comment/entity"
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	ID        string
	PhotoURL  string
	Title     string
	Caption   string
	UserID    string
	Comments  []ec.Comment
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
