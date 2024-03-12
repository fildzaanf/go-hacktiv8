package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (order *Order) BeforeCreate(tx *gorm.DB) (err error) {
	UUID := uuid.New()
	order.ID = UUID.String()
	return nil
}

func (item *Item) BeforeCreate(tx *gorm.DB) (err error) {
	UUID := uuid.New()
	item.ID = UUID.String()
	return nil
}
