package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	UUID := uuid.New()
	user.ID = UUID.String()

	if user.Role == "" {
		user.Role = "user" 
	}
	
	return nil
}

