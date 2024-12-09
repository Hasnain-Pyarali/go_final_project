package config

import (
    "final/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    // Auto-migrate tables
    database.AutoMigrate(&models.User{}, &models.Task{})

    DB = database
}
