package database

import (
	"log"

	"technical-assignment/backend/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	err = database.AutoMigrate(&models.Person{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = database
}
