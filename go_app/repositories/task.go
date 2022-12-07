package repositories

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zimnushka/task_me_go/go_app/models"
)

type TaskRepository struct {
	taskMeDB TaskMeDB
}

func (taskRepository TaskRepository) GetTaskFromId(id int) (*models.Task, error) {
	db, err := taskRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT * FROM tasks WHERE id = '%d' LIMIT 1", id)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	for results.Next() {
		var item models.Task
		err := results.Scan(&item.Id, &item.Title, &item.Description, &item.Time, &item.ProjectId, &item.Status, &item.UserId, &item.Cost)
		if err != nil {
			return nil, err
		}
		return &item, nil
	}

	return nil, errors.New("Unexpected error user repository")
}

func (taskRepository TaskRepository) GetTasks() ([]models.Task, error) {
	db, err := taskRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	results, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer results.Close()

	itemsLng := 0
	items := make([]models.Task, itemsLng)

	for results.Next() {
		var item models.Task
		err := results.Scan(&item.Id, &item.Title, &item.Description, &item.Time, &item.ProjectId, &item.Status, &item.UserId, &item.Cost)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
		itemsLng++
	}

	return items, nil
}

func (taskRepository TaskRepository) GetTasksFromUser(user_id int) ([]models.Task, error) {
	db, err := taskRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT * FROM tasks WHERE user_id = '%d'", user_id)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	itemsLng := 0
	items := make([]models.Task, itemsLng)

	for results.Next() {
		var item models.Task
		err := results.Scan(&item.Id, &item.Title, &item.Description, &item.Time, &item.ProjectId, &item.Status, &item.UserId, &item.Cost)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
		itemsLng++
	}

	return items, nil
}

func (taskRepository TaskRepository) GetTasksFromProject(projectId int) ([]models.Task, error) {
	db, err := taskRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT * FROM tasks WHERE project_id = '%d'", projectId)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	itemsLng := 0
	items := make([]models.Task, itemsLng)

	for results.Next() {
		var item models.Task
		err := results.Scan(&item.Id, &item.Title, &item.Description, &item.Time, &item.ProjectId, &item.Status, &item.UserId, &item.Cost)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
		itemsLng++
	}

	return items, nil
}

func (taskRepository TaskRepository) AddTask(task models.Task) (*models.Task, error) {
	db, err := taskRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	userIdLabel := "NULL"
	if task.UserId != nil {
		userIdLabel = fmt.Sprint("'", *task.UserId, "'")
	}
	query := fmt.Sprint("INSERT INTO tasks (title, description, project_id, due_date, status_id, user_id, cost) VALUES ('", task.Title, "','", task.Description, "','", task.ProjectId, "','", task.Time, "','", task.Status, "',", userIdLabel, ",'", task.Cost, "') RETURNING id")
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	for results.Next() {
		err = results.Scan(&task.Id)
	}
	return &task, err
}

func (taskRepository TaskRepository) UpdateTask(task models.Task) error {
	db, err := taskRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return err
	}
	userIdLabel := "NULL"
	if task.UserId != nil {
		userIdLabel = fmt.Sprint("'", *task.UserId, "'")
	}
	query := fmt.Sprintf("UPDATE tasks SET title = '%s', description = '%s', project_id = '%d', due_date = '%s', status_id = '%d', user_id = %s, cost = '%d' WHERE id = %d", task.Title, task.Description, task.ProjectId, task.Time, task.Status, userIdLabel, task.Cost, *task.Id)
	results, err := db.Query(query)
	if err == nil {
		defer results.Close()
	}

	return err
}

func (taskRepository TaskRepository) DeleteTask(id int) error {
	db, err := taskRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM tasks WHERE id = %d", id)
	results, err := db.Query(query)
	if err != nil {
		return err
	}
	defer results.Close()
	return nil

}
