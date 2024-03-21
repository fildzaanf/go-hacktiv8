package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        string `gorm:"primarykey"`
	UserID    string `gorm:"not null"`
	PhotoID   string `gorm:"not null"`
	Message   string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
