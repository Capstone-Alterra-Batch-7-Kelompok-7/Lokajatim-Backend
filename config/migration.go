package config

import (
	"lokajatim/entities"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&entities.User{})

}
