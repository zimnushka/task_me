package usecases

import (
	"errors"
	"time"

	"github.com/zimnushka/task_me_go/go_app/models"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

type TimeIntervalUseCase struct {
	intervalRepository repositories.IntervalRepository

	taskUseCase TaskUseCase
}

func (useCase *TimeIntervalUseCase) AddInterval(taskId, userId int) (*models.Interval, error) {
	access, err := useCase.taskUseCase.CheckUserHaveTask(taskId, userId)
	if err != nil {
		return nil, err
	}
	if access {
		var item models.Interval
		item.TaskId = taskId
		item.UserId = userId
		item.TimeStart = time.Now().Format(time.RFC3339)
		return useCase.intervalRepository.Add(item)
	}
	return nil, errors.New("Forbiden")
}

func (useCase *TimeIntervalUseCase) GetIntervalById(id, userId int) (*models.Interval, error) {
	interval, err := useCase.intervalRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	access, err := useCase.taskUseCase.CheckUserHaveTask(interval.TaskId, userId)
	if err != nil {
		return nil, err
	}
	if access {
		return interval, err
	}
	return nil, errors.New("Forbiden")
}

func (useCase *TimeIntervalUseCase) GetIntervalsByTask(id, userId int) ([]models.Interval, error) {
	access, err := useCase.taskUseCase.CheckUserHaveTask(id, userId)
	if err != nil {
		return nil, err
	}
	if access {
		return useCase.intervalRepository.GetByTaskId(id)
	}
	return nil, errors.New("Forbiden")
}

func (useCase *TimeIntervalUseCase) GetIntervalsByUser(userId int) ([]models.Interval, error) {
	return useCase.intervalRepository.GetByUserId(userId)
}

func (useCase *TimeIntervalUseCase) UpdateInterval(item models.Interval, userId int) error {
	access, err := useCase.taskUseCase.CheckUserHaveTask(item.TaskId, userId)
	if err != nil {
		return err
	}
	if access {
		return useCase.intervalRepository.Update(item)
	}
	return errors.New("Forbiden")
}

func (useCase *TimeIntervalUseCase) FinishInterval(taskId, userId int) error {
	access, err := useCase.taskUseCase.CheckUserHaveTask(taskId, userId)
	if err != nil {
		return err
	}
	if access {
		item, err := useCase.intervalRepository.GetNotEndedInterval(taskId, userId)
		if err != nil {
			return err
		}
		item.TimeEnd = time.Now().Format(time.RFC3339)
		return useCase.intervalRepository.Update(*item)
	}
	return errors.New("Forbiden")
}
