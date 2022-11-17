package models

import "time"

type Task struct {
	Id          *int
	ProjectId   int
	Title       string
	Description string
	Time        time.Time
}

type TaskUser struct {
	UserId int
	TaskId int
}
