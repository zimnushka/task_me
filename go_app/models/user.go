package models

type User struct {
	Id       *int   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Color    int    `json:"color"`
	Cost     int    `json:"cost"`
}

func (item User) ToDTO() UserDTO {
	return UserDTO{
		Id:    *item.Id,
		Name:  item.Name,
		Email: item.Email,
		Color: item.Color,
	}
}

type UserDTO struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Color int    `json:"color"`
}
