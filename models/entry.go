package models

import (
	"gorm.io/gorm"
)

type EntryType string
type ActionType string

const (
	PostAction  ActionType = "write"
	WatchAction ActionType = "watch"
	ReadAction  ActionType = "read"
	PostEntry   EntryType  = "post"
)

type Entry struct {
	gorm.Model
	Type      EntryType
	Title     string
	Path      string
	SubjectID int
	Subject   EntrySubject // Foreign id
}
