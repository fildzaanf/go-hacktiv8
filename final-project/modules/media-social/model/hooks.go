package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (medsos *MediaSocial) BeforeCreate(tx *gorm.DB) (err error) {
	UUID := uuid.New()
	medsos.ID = UUID.String()
	return nil
}
