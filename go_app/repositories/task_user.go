package repositories

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zimnushka/task_me_go/go_app/models"
)

type TaskUserRepository struct {
	taskMeDB TaskMeDB
}

func (repository TaskUserRepository) GetTasksByUser(id int) ([]models.Task, error) {
	db, err := repository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT tasks.id, tasks.title, tasks.description, tasks.due_date, tasks.project_id FROM UsersTasks INNER JOIN tasks ON UsersTasks.task_id=tasks.id AND UsersTasks.user_id=%d", id)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	itemsLng := 0
	items := make([]models.Task, itemsLng)

	for results.Next() {
		var item models.Task
		err := results.Scan(&item.Id, &item.Title, &item.Description, &item.Time, &item.ProjectId)
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
	query := fmt.Sprintf("SELECT users.id, users.name, users.email FROM UsersTasks INNER JOIN users ON UsersTasks.user_id=users.id AND UsersTasks.task_id=%d", id)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	itemsLng := 0
	items := make([]models.User, itemsLng)

	for results.Next() {
		var item models.User
		err := results.Scan(&item.Id, &item.Name, &item.Email)
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
	query := fmt.Sprintf("INSERT INTO UsersTasks (user_id, task_id) VALUES ('%d','%d')", userId, taskId)
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

	query := fmt.Sprintf("DELETE FROM UsersTasks WHERE user_id = %d AND task_id = %d", userId, taskId)
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

	query := fmt.Sprintf("DELETE FROM UsersTasks WHERE task_id = %d", taskId)
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

	query := fmt.Sprintf("DELETE FROM UsersTasks WHERE user_id = %d", userId)
	results, err := db.Query(query)
	if err == nil {
		defer results.Close()
	}

	return err

}
