package models

import "time"

type Task struct {
	Id          *int      `json:"id"`
	ProjectId   int       `json:"projectId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Time        time.Time `json:"dueDate"`
}

type TaskUser struct {
	UserId int
	TaskId int
}
