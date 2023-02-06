package usecases

import (
	"errors"

	"github.com/zimnushka/task_me_go/go_app/models"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

type TimeIntervalUseCase struct {
	intervalRepository repositories.IntervalRepository

	taskUseCase TaskUseCase
}

func (useCase *TimeIntervalUseCase) AddInterval(item models.Interval, userId int) (*models.Interval, error) {
	access, err := useCase.taskUseCase.CheckUserHaveTask(item.TaskId, userId)
	if err != nil {
		return nil, err
	}
	if access {
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

func (useCase *TimeIntervalUseCase) GetNotEndedInterval(userId int) (*models.Interval, error) {
	interval, err := useCase.intervalRepository.GetNotEndedIntervalByUserId(userId)
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
	access, err := useCase.taskUseCase.CheckUserHaveTask(*&item.TaskId, userId)
	if err != nil {
		return err
	}
	if access {
		return useCase.intervalRepository.Update(item)
	}
	return errors.New("Forbiden")
}
