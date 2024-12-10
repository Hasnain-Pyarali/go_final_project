package config

import (
<<<<<<< HEAD
    "final/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
=======
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"final/models"
>>>>>>> origin/main
)

var DB *gorm.DB

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed")
}

func ConnectDatabase() {
<<<<<<< HEAD
    database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    // Auto-migrate tables
    database.AutoMigrate(&models.User{}, &models.Task{})

    DB = database
=======
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	DB = database
	Migrate(DB)
>>>>>>> origin/main
}
