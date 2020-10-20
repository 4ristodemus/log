package db

import (
	"github.com/4ristodemus/log/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func InitDatabase() {
	db, err := gorm.Open(sqlite.Open("data/test.db"), &gorm.Config{})
	Database = db

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate
	Database.AutoMigrate(&models.EntrySubject{}, &models.Entry{})
}
