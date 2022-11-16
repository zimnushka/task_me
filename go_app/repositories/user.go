package repositories

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zimnushka/task_me_go/go_app/models"
)

type UserRepository struct {
	taskMeDB TaskMeDB
}

func (userRepository UserRepository) GetUserFromEmail(email string) (*models.User, error) {
	db, err := userRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT * FROM users WHERE email = '%s' LIMIT 1", email)
	results, err := db.Query(query)
	defer results.Close()
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var user models.User
		err := results.Scan(&user.Id, &user.Name, &user.Password, &user.Email)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}

	return nil, errors.New("Unexpected error user repository")
}

func (userRepository UserRepository) GetUserFromId(id int) (*models.User, error) {
	db, err := userRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT * FROM users WHERE id = '%d' LIMIT 1", id)
	results, err := db.Query(query)
	defer results.Close()
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var user models.User
		err := results.Scan(&user.Id, &user.Name, &user.Password, &user.Email)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}

	return nil, errors.New("Unexpected error user repository")
}

func (userRepository UserRepository) GetUsers() ([]models.User, error) {
	db, err := userRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	results, err := db.Query("SELECT * FROM users")
	defer results.Close()
	if err != nil {
		return nil, err
	}

	usersLng := 0
	users := make([]models.User, usersLng)

	for results.Next() {
		var user models.User
		err := results.Scan(&user.Id, &user.Name, &user.Password, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
		usersLng++
	}

	return users, nil
}

func (userRepository UserRepository) AddUser(user models.User) (*models.User, error) {
	db, err := userRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("INSERT INTO users (name, password, email) VALUES ('%s','%s','%s') RETURNING id", user.Name, user.Password, user.Email)
	results, err := db.Query(query)
	defer results.Close()

	for results.Next() {
		err = results.Scan(&user.Id)
	}
	return &user, err
}

func (userRepository UserRepository) UpdateUser(user models.User) error {
	db, err := userRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("UPDATE users SET name = '%s', password = '%s', email = '%s' WHERE id = %d", user.Name, user.Password, user.Email, *user.Id)
	results, err := db.Query(query)
	defer results.Close()

	return err
}

func (userRepository UserRepository) DeleteUser(id int) error {
	db, err := userRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM users WHERE id = %d", id)
	results, err := db.Query(query)
	defer results.Close()
	if err != nil {
		return err
	}
	return nil

}
