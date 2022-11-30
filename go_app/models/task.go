package models

import "time"

const (
	StatusOpen       = 0
	StatusInProgress = 1
	StatusReview     = 2
	StatusClose      = 3
)

type Task struct {
	Id          *int      `json:"id"`
	ProjectId   int       `json:"projectId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      int       `json:"StatusId"`
	Time        time.Time `json:"dueDate"`
}

type TaskUser struct {
	UserId int
	TaskId int
}
