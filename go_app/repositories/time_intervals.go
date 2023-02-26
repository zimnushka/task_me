package repositories

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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
	query := fmt.Sprintf("SELECT * FROM intervals WHERE id = '%d'", id)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	for results.Next() {
		var item models.Interval
		err := results.Scan(&item.Id, &item.TaskId, &item.TimeStart, &item.TimeEnd, &item.UserId)
		if err != nil {
			return nil, err
		}

		return &item, err
	}

	return nil, errors.New("Unexpected error user repository")
}

func (intervalRepository IntervalRepository) GetNotEndedInterval(taskId, userId int) (*models.Interval, error) {
	db, err := intervalRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT * FROM intervals WHERE user_id = '%d' AND task_id = '%d' AND time_end = ''", userId, taskId)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	for results.Next() {
		var item models.Interval
		err := results.Scan(&item.Id, &item.TaskId, &item.TimeStart, &item.TimeEnd, &item.UserId)
		if err != nil {
			return nil, err
		}

		return &item, err
	}

	return nil, errors.New("Unexpected error user repository")
}

func (intervalRepository IntervalRepository) GetByTaskId(id int) ([]models.Interval, error) {
	db, err := intervalRepository.taskMeDB.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT * FROM intervals WHERE task_id = '%d'", id)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	itemsLng := 0
	items := make([]models.Interval, itemsLng)

	for results.Next() {
		var item models.Interval
		err := results.Scan(&item.Id, &item.TaskId, &item.TimeStart, &item.TimeEnd, &item.UserId)
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
	query := fmt.Sprintf("SELECT * FROM intervals WHERE user_id = '%d'", id)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	itemsLng := 0
	items := make([]models.Interval, itemsLng)

	for results.Next() {
		var item models.Interval
		err := results.Scan(&item.Id, &item.TaskId, &item.TimeStart, &item.TimeEnd, &item.UserId)
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

	query := fmt.Sprintf("INSERT INTO intervals (task_id, user_id, time_start, time_end) VALUES ('%d','%d','%s','%s') RETURNING id", item.TaskId, item.UserId, item.TimeStart, item.TimeEnd)
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
	query := fmt.Sprintf("UPDATE intervals SET task_id = '%d', user_id = '%d', time_start = '%s', time_end = '%s' WHERE id = %d", item.TaskId, item.UserId, item.TimeStart, item.TimeEnd, id)
	results, err := db.Query(query)
	if err == nil {
		defer results.Close()
	}

	return err
}
