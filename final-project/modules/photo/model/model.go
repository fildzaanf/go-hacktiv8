package model

import (
	cm "final-project/modules/comment/model"
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	ID        string       `gorm:"primarykey"`
	PhotoURL  string       `gorm:"not null"`
	Title     string       `gorm:"not null"`
	Caption   string       `gorm:"not null"`
	UserID    string       `gorm:"not null"`
	Comments  []cm.Comment `gorm:"foreignKey:PhotoID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
