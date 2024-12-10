package config

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"final/models"
)

var DB *gorm.DB

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.Task{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed")
}

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	Migrate(database)
	DB = database
}
