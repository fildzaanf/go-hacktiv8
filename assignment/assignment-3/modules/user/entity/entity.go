package entity

import (
	"assignment-2/modules/order/entity"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string
	Fullname  string
	Email     string
	Password  string
	Orders    []entity.Order
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
