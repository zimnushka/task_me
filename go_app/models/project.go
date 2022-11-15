package models

type Project struct {
	Id    *int
	Title string
	Color int
}

type UserProject struct {
	UserId    int
	ProjectId int
}
