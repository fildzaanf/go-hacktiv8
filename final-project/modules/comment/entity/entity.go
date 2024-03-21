package entity

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        string
	UserID    string
	PhotoID   string
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
