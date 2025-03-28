package dto

import "time"

type Task struct {
	ID       uint      `json:"task_id"`
	OwnerID  uint      `json:"owner_id"`
	Title    string    `json:"title"`
	Deadline time.Time `json:"deadline"`
	Status   string    `json:"status"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}
