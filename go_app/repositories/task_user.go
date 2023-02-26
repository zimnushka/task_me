package repositories

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zimnushka/task_me_go/go_app/models"
)

type TaskUserRepository struct {
	taskMeDB TaskMeDB
}

func (taskRepository TaskUserRepository) GetTasksByUser(id int) ([]models.Task, error) {
	db, err := taskRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT tasks.id,  tasks.title, tasks.description, tasks.start_date, tasks.stop_date, tasks.project_id, tasks.cost, tasks.status_id FROM TasksUsers INNER JOIN tasks ON TasksUsers.task_id=tasks.id AND TasksUsers.user_id=%d", id)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	itemsLng := 0
	items := make([]models.Task, itemsLng)

	for results.Next() {
		var item models.Task
		err := results.Scan(&item.Id, &item.Title, &item.Description, &item.StartDate, &item.StopDate, &item.ProjectId, &item.Cost, &item.Status)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
		itemsLng++
	}

	return items, nil
}

func (repository TaskUserRepository) GetUsersByTask(id int) ([]models.User, error) {
	db, err := repository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT users.id, users.name, users.email, users.color, users.cost FROM TasksUsers INNER JOIN users ON TasksUsers.user_id=users.id AND TasksUsers.task_id=%d", id)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	itemsLng := 0
	items := make([]models.User, itemsLng)

	for results.Next() {
		var item models.User
		err := results.Scan(&item.Id, &item.Name, &item.Email, &item.Color, &item.Cost)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
		itemsLng++
	}

	return items, nil
}

func (repository TaskUserRepository) AddLink(taskId, userId int) error {
	db, err := repository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("INSERT INTO TasksUsers (user_id, task_id) VALUES ('%d','%d')", userId, taskId)
	results, err := db.Query(query)
	if err == nil {
		defer results.Close()
	}

	return err
}

func (repository TaskUserRepository) DeleteLink(taskId, userId int) error {
	db, err := repository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM TasksUsers WHERE user_id = %d AND task_id = %d", userId, taskId)
	results, err := db.Query(query)
	if err == nil {
		defer results.Close()
	}

	return err
}

func (repository TaskUserRepository) DeleteAllLinkByTask(taskId int) error {
	db, err := repository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM TasksUsers WHERE task_id = %d", taskId)
	results, err := db.Query(query)
	if err == nil {
		defer results.Close()
	}

	return err
}

func (repository TaskUserRepository) DeleteAllLinkByUser(userId int) error {
	db, err := repository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM TasksUsers WHERE user_id = %d", userId)
	results, err := db.Query(query)
	if err == nil {
		defer results.Close()
	}

	return err

}
