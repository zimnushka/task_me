package models

type Project struct {
	Id    *int
	Title string
	Color int
}

type ProjectUser struct {
	UserId    int
	ProjectId int
}
