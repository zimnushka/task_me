package repositories

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/zimnushka/task_me_go/go_app/models"
)

func GetUsers() []models.User {
	db := getDB()
	defer db.Close()
	results := queryDB(db, "SELECT * FROM users")
	defer results.Close()

	usersLng := 0
	users := make([]models.User, usersLng)

	for results.Next() {
		var user models.User
		err := results.Scan(&user.Id, &user.Name, &user.Password, &user.Email)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
		usersLng++
	}

	return users
}
