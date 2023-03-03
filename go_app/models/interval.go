package models

type Interval struct {
	Id          *int   `json:"id"`
	TaskId      int    `json:"task_id"`
	UserId      int    `json:"user_id"`
	TimeStart   string `json:"time_start"`
	TimeEnd     string `json:"time_end"`
	Description string `json:"description"`
}
