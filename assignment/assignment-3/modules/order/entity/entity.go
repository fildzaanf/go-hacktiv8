package entity

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID           string
	CustomerName string
	OrderedAt    time.Time
	Items        []Item
	UserID       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

type Item struct {
	ID          string
	ItemCode    string
	Description string
	Quantity    int
	OrderID     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
