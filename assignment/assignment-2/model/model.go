package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID           string    `gorm:"primaryKey"`
	CustomerName string    `gorm:"not null"`
	OrderedAt    time.Time `gorm:"type:timestamp;default:now()"`
	Items        []Item    `gorm:"foreignKey:OrderID;references:ID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Item struct {
	ID          string `gorm:"primaryKey"`
	ItemCode    string `gorm:"not null"`
	Description string `gorm:"not null"`
	Quantity    int    `gorm:"not null"`
	OrderID     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
