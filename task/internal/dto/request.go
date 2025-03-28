package dto

import "time"

type CreateTaskRequest struct {
	OwnerID  uint
	Title    string
	Deadline time.Time
	Status   string
}

type CompleteTaskRequest struct {
	TaskID  uint
	OwnerID uint
}

type DeleteTaskRequest struct {
	TaskID  uint
	OwnerID uint
}
