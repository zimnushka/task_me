package repositories

import (
	"fmt"

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

func AddUser(user models.User) {
	db := getDB()
	defer db.Close()
	query := fmt.Sprintf("INSERT INTO users (name, password, email) VALUES ('%s','%s','%s')", user.Name, user.Password, user.Email)
	results := queryDB(db, query)
	defer results.Close()

}

func UpdateUser(user models.User) {
	db := getDB()
	defer db.Close()
	query := fmt.Sprintf("UPDATE users SET name = '%s', password = '%s', email = '%s' WHERE id = %d", user.Name, user.Password, user.Email, *user.Id)
	results := queryDB(db, query)
	defer results.Close()

}

func DeleteUser(id int) {
	db := getDB()
	defer db.Close()
	query := fmt.Sprintf("DELETE FROM users WHERE id = %d", id)
	results := queryDB(db, query)
	defer results.Close()

}
