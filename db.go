package main

import (
	"github.com/4ristodemus/logs/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() {
	db, err := gorm.Open(sqlite.Open("data/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate
	db.AutoMigrate(&models.EntrySubject{}, &models.Entry{})
}
