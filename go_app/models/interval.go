package models

type Interval struct {
	Id          *int    `json:"id"`
	Task        TaskDTO `json:"task"`
	User        UserDTO `json:"user"`
	TimeStart   string  `json:"time_start"`
	TimeEnd     string  `json:"time_end"`
	Description string  `json:"description"`
}
