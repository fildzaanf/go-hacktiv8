package migration

import (
	"assignment-2/modules/order/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.Order{},
		&model.Item{},
	)
}
