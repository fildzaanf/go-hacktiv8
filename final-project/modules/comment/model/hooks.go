package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (comment *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	UUID := uuid.New()
	comment.ID = UUID.String()
	return nil
}
