package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (photo *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	UUID := uuid.New()
	photo.ID = UUID.String()
	return nil
}
