package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	OwnerID  uint
	Title    string
	Deadline time.Time
	Status   string
}
