package model

import (
	"assignment-2/modules/order/model"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primarykey"`
	Fullname  string
	Email     string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Orders    []model.Order  `gorm:"foreignKey:UserID;references:ID"`
}
