package config

import (
	"lokajatim/entities"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(
		&entities.User{},
		&entities.Event{},
		&entities.Ticket{},
		&entities.Article{},
		&entities.Comment{},
		&entities.Like{},
		&entities.Category{},
		&entities.Product{},
		&entities.ProductPhoto{},
	)

}
