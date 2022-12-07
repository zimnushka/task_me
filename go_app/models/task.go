package models

const (
	StatusOpen       = 0
	StatusInProgress = 1
	StatusReview     = 2
	StatusClose      = 3
)

type Task struct {
	Id          *int   `json:"id"`
	ProjectId   int    `json:"projectId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      int    `json:"statusId"`
	Time        string `json:"dueDate"`
	UserId      *int   `json:"userId"`
	Cost        int    `json:"cost"`
}
