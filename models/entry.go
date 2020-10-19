package models

import (
	"gorm.io/gorm"
)

type ActionType string

const (
	PostAction  ActionType = "write"
	WatchAction            = "watch"
	ReadAction             = "read"
)

type Entry struct {
	gorm.Model
	Type    ActionType
	Title   string
	Subject EntrySubject // Foreign id
}
