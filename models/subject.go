package models

import (
	"gorm.io/gorm"
)

type EntrySubject struct {
	gorm.Model
	ID       int
	Title    string
	Metadata string
}
