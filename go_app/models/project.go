package models

type Project struct {
	Id    *int   `json:"id"`
	Title string `json:"title"`
	Color int    `json:"color"`
}

type ProjectUser struct {
	UserId    int
	ProjectId int
}
