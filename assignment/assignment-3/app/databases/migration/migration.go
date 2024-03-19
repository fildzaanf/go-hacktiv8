package migration

import (
	om "assignment-2/modules/order/model"
	um "assignment-2/modules/user/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&um.User{},
		&om.Order{},
		&om.Item{},
	)
}
