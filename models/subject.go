package models

import (
	"gorm.io/gorm"
)

type EntrySubject struct {
	gorm.Model
	Title    string
	Metadata string
}
