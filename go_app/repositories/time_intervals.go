package repositories

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zimnushka/task_me_go/go_app/app_errors"
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
	query := fmt.Sprintf("SELECT intervals.id, intervals.task_id, intervals.time_start, intervals.time_end, intervals.user_id, description FROM intervals INNER JOIN users ON intervals.user_id = users.id AND intervals.id = '%d'", id)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	for results.Next() {
		var item models.Interval
		err := results.Scan(&item.Id, &item.TaskId, &item.TimeStart, &item.TimeEnd, &item.User.Id, &item.Description, &item.User.Name, &item.User.Email, &item.User.Color)
		if err != nil {
			return nil, err
		}

		return &item, err
	}

	return nil, errors.New(app_errors.ERR_Unexpected_repository_error)
}

func (intervalRepository IntervalRepository) GetNotEndedInterval(taskId, userId int) (*models.Interval, error) {
	db, err := intervalRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT intervals.id, intervals.task_id, intervals.time_start, intervals.time_end, intervals.user_id, intervals.description, users.name, users.email, users.color FROM intervals INNER JOIN users ON intervals.user_id = users.id AND intervals.user_id = '%d' AND intervals.task_id = '%d' AND intervals.time_end = ''", userId, taskId)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	for results.Next() {
		var item models.Interval
		err := results.Scan(&item.Id, &item.TaskId, &item.TimeStart, &item.TimeEnd, &item.User.Id, &item.Description, &item.User.Name, &item.User.Email, &item.User.Color)
		if err != nil {
			return nil, err
		}

		return &item, err
	}

	return nil, errors.New(app_errors.ERR_Unexpected_repository_error)
}

func (intervalRepository IntervalRepository) GetByTaskId(id int) ([]models.Interval, error) {
	db, err := intervalRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT intervals.id, intervals.task_id, intervals.time_start, intervals.time_end, intervals.user_id, intervals.description, users.name, users.email, users.color FROM intervals INNER JOIN users ON intervals.user_id = users.id AND intervals.task_id = '%d'", id)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	itemsLng := 0
	items := make([]models.Interval, itemsLng)

	for results.Next() {
		var item models.Interval
		err := results.Scan(&item.Id, &item.TaskId, &item.TimeStart, &item.TimeEnd, &item.User.Id, &item.Description, &item.User.Name, &item.User.Email, &item.User.Color)
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
	query := fmt.Sprintf("SELECT intervals.id, intervals.task_id, intervals.time_start, intervals.time_end, intervals.user_id, intervals.description, users.name, users.email, users.color FROM intervals INNER JOIN users ON intervals.user_id = users.id AND intervals.user_id = '%d'", id)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	itemsLng := 0
	items := make([]models.Interval, itemsLng)

	for results.Next() {
		var item models.Interval
		err := results.Scan(&item.Id, &item.TaskId, &item.TimeStart, &item.TimeEnd, &item.User.Id, &item.Description, &item.User.Name, &item.User.Email, &item.User.Color)
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
	defer db.Close()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("INSERT INTO intervals (task_id, user_id, time_start, time_end, description) VALUES ('%d','%d','%s','%s','%s') RETURNING id", item.TaskId, item.User.Id, item.TimeStart, item.TimeEnd, item.Description)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	for results.Next() {
		err = results.Scan(&item.Id)
	}
	return &item, err
}

func (intervalRepository IntervalRepository) Update(item models.Interval) error {
	db, err := intervalRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return err
	}
	var id int
	id = *item.Id
	query := fmt.Sprintf("UPDATE intervals SET task_id = '%d', user_id = '%d', time_start = '%s', time_end = '%s', description = '%s' WHERE id = %d", item.TaskId, item.User.Id, item.TimeStart, item.TimeEnd, item.Description, id)
	results, err := db.Query(query)
	if err == nil {
		defer results.Close()
	}

	return err
}
