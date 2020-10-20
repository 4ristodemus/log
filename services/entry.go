package services

import (
	"github.com/4ristodemus/log/db"
	"github.com/4ristodemus/log/models"
)

func AllEntries() (err error, entries []models.Entry) {
	err = db.Database.Order("created_at DESC").Find(&entries).Error
	return err, entries
}
