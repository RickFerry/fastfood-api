package database

import (
	"fastfood-api/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("fastfood.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	err = db.AutoMigrate(&models.MenuItem{})
	if err != nil {
		log.Fatal("Failed to migrate database")
	}

	return db
}