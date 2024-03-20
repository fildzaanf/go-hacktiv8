package migration

import (
	"assignment-2/model"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Order{},
		&model.Item{},
	)

	migrator := db.Migrator()

	tables := []string{"users", "orders", "items"}

	for _, table := range tables {
		if !migrator.HasTable(table) {
			log.Fatalf("table %s was not successfully created", table)
		}
	}

	log.Println("all tables were successfully migrated")
}
