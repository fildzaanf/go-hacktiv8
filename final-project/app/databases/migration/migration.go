package migration

import (
	cm "final-project/modules/comment/model"
	msm "final-project/modules/media-social/model"
	pm "final-project/modules/photo/model"
	um "final-project/modules/user/model"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&um.User{},
		&pm.Photo{},
		&cm.Comment{},
		&msm.MediaSocial{},
	)

	migrator := db.Migrator()
	tables := []string{"users", "photos", "comments", "media_socials"}
	for _, table := range tables {
		if !migrator.HasTable(table) {
			log.Fatalf("table %s was not successfully created", table)
		}
	}
	log.Println("all tables were successfully migrated")
}
