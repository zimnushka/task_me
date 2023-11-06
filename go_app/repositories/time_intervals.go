package repositories

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zimnushka/task_me_go/go_app/app"
	"github.com/zimnushka/task_me_go/go_app/models"
)

type IntervalRepository struct {
	taskMeDB TaskMeDB
}

func (intervalRepository IntervalRepository) GetById(id int) (*models.Interval, error) {
	db, err := intervalRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT intervals.id, intervals.time_start, intervals.time_end, intervals.user_id, intervals.description, users.name, users.email, users.color, tasks.id, tasks.project_id, tasks.title, tasks.status_id, tasks.start_date, tasks.stop_date, tasks.cost FROM intervals INNER JOIN users ON intervals.user_id = users.id INNER JOIN tasks ON intervals.task_id = tasks.id AND intervals.id = '%d'", id)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	for results.Next() {
		var item models.Interval
		err := results.Scan(&item.Id, &item.TimeStart, &item.TimeEnd, &item.User.Id, &item.Description, &item.User.Name, &item.User.Email, &item.User.Color, &item.Task.Id, &item.Task.ProjectId, &item.Task.Title, &item.Task.Status, &item.Task.StartDate, &item.Task.StopDate, &item.Task.Cost)
		if err != nil {
			return nil, err
		}

		return &item, err
	}

	return nil, errors.New(app.ERR_Unexpected_repository_error)
}

func (intervalRepository IntervalRepository) GetNotEndedInterval(userId int) (*models.Interval, error) {
	db, err := intervalRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT intervals.id, intervals.time_start, intervals.time_end, intervals.user_id, intervals.description, users.name, users.email, users.color, tasks.id, tasks.project_id, tasks.title, tasks.status_id, tasks.start_date, tasks.stop_date, tasks.cost FROM intervals INNER JOIN users ON intervals.user_id = users.id INNER JOIN tasks ON intervals.task_id = tasks.id AND intervals.user_id = '%d' AND intervals.time_end = ''", userId)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	for results.Next() {
		var item models.Interval
		err := results.Scan(&item.Id, &item.TimeStart, &item.TimeEnd, &item.User.Id, &item.Description, &item.User.Name, &item.User.Email, &item.User.Color, &item.Task.Id, &item.Task.ProjectId, &item.Task.Title, &item.Task.Status, &item.Task.StartDate, &item.Task.StopDate, &item.Task.Cost)
		if err != nil {
			return nil, err
		}

		return &item, err
	}

	return nil, errors.New(app.ERR_Unexpected_repository_error)
}

func (intervalRepository IntervalRepository) GetByTaskId(id int) ([]models.Interval, error) {
	db, err := intervalRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT intervals.id, intervals.time_start, intervals.time_end, intervals.user_id, intervals.description, users.name, users.email, users.color, tasks.id, tasks.project_id, tasks.title, tasks.status_id, tasks.start_date, tasks.stop_date, tasks.cost FROM intervals INNER JOIN users ON intervals.user_id = users.id INNER JOIN tasks ON intervals.task_id = tasks.id AND intervals.task_id = '%d'", id)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	itemsLng := 0
	items := make([]models.Interval, itemsLng)

	for results.Next() {
		var item models.Interval
		err := results.Scan(&item.Id, &item.TimeStart, &item.TimeEnd, &item.User.Id, &item.Description, &item.User.Name, &item.User.Email, &item.User.Color, &item.Task.Id, &item.Task.ProjectId, &item.Task.Title, &item.Task.Status, &item.Task.StartDate, &item.Task.StopDate, &item.Task.Cost)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
		itemsLng++
	}

	return items, nil
}

func (intervalRepository IntervalRepository) GetByProjectId(id int) ([]models.Interval, error) {
	db, err := intervalRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT intervals.id, intervals.time_start, intervals.time_end, intervals.user_id, intervals.description, users.name, users.email, users.color, tasks.id, tasks.project_id, tasks.title, tasks.status_id, tasks.start_date, tasks.stop_date, tasks.cost FROM intervals INNER JOIN users ON intervals.user_id = users.id INNER JOIN tasks ON intervals.task_id = tasks.id INNER JOIN projects ON tasks.project_id = projects.id AND projects.id = '%d'", id)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	itemsLng := 0
	items := make([]models.Interval, itemsLng)

	for results.Next() {
		var item models.Interval
		err := results.Scan(&item.Id, &item.TimeStart, &item.TimeEnd, &item.User.Id, &item.Description, &item.User.Name, &item.User.Email, &item.User.Color, &item.Task.Id, &item.Task.ProjectId, &item.Task.Title, &item.Task.Status, &item.Task.StartDate, &item.Task.StopDate, &item.Task.Cost)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
		itemsLng++
	}

	return items, nil
}

func (intervalRepository IntervalRepository) GetByUserId(id int) ([]models.Interval, error) {
	db, err := intervalRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT intervals.id, intervals.time_start, intervals.time_end, intervals.user_id, intervals.description, users.name, users.email, users.color, tasks.id, tasks.project_id, tasks.title, tasks.status_id, tasks.start_date, tasks.stop_date, tasks.cost FROM intervals INNER JOIN users ON intervals.user_id = users.id INNER JOIN tasks ON intervals.task_id = tasks.id AND intervals.user_id = '%d'", id)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	itemsLng := 0
	items := make([]models.Interval, itemsLng)

	for results.Next() {
		var item models.Interval
		err := results.Scan(&item.Id, &item.TimeStart, &item.TimeEnd, &item.User.Id, &item.Description, &item.User.Name, &item.User.Email, &item.User.Color, &item.Task.Id, &item.Task.ProjectId, &item.Task.Title, &item.Task.Status, &item.Task.StartDate, &item.Task.StopDate, &item.Task.Cost)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
		itemsLng++
	}

	return items, nil
}

func (intervalRepository IntervalRepository) Add(item models.Interval) (*models.Interval, error) {
	db, err := intervalRepository.taskMeDB.GetDB()
	if err != nil {
		defer db.Close()
		return nil, err
	}
	defer db.Close()
	query := "INSERT INTO intervals (task_id, user_id, time_start, time_end, description) VALUES (95,1,'2023-11-06T22:49:54+03:00','','') RETURNING id"
	// query := fmt.Sprintf("INSERT INTO intervals (task_id, user_id, time_start, time_end, description) VALUES (%d,%d,'%s','%s','%s') RETURNING id", *item.Task.Id, item.User.Id, item.TimeStart, item.TimeEnd, item.Description)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		err = results.Scan(&item.Id)
	}
	defer results.Close()
	return &item, err
}

func (intervalRepository IntervalRepository) Update(item models.Interval) error {
	db, err := intervalRepository.taskMeDB.GetDB()
	if err != nil {
		defer db.Close()
		return err
	}
	defer db.Close()
	query := fmt.Sprintf("UPDATE intervals SET task_id = %d, user_id = %d, time_start = '%s', time_end = '%s', description = '%s' WHERE id = %d", *item.Task.Id, *item.User.Id, item.TimeStart, item.TimeEnd, item.Description, *item.Id)
	results, err := db.Query(query)
	if err == nil {
		defer results.Close()
	}

	return err
}
