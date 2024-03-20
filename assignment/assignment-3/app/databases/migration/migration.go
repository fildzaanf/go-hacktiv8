package migration

import (
	om "assignment-2/modules/order/model"
	um "assignment-2/modules/user/model"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&um.User{})
	db.AutoMigrate(&om.Order{}, &om.Item{})

	migrator := db.Migrator()
	tables := []string{"users", "orders", "items"}
	for _, table := range tables {
		if !migrator.HasTable(table) {
			log.Fatalf("table %s was not successfully created", table)
		}
	}
	log.Println("all tables were successfully migrated")
}
