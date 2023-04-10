package usecases

import (
	"time"

	"github.com/zimnushka/task_me_go/go_app/models"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

type TimeIntervalUseCase struct {
	intervalRepository repositories.IntervalRepository
	userRepository     repositories.UserRepository

	taskUseCase TaskUseCase
}

func (useCase *TimeIntervalUseCase) AddInterval(taskId, userId int) (*models.Interval, error) {
	if err := useCase.taskUseCase.CheckUserHaveTask(taskId, userId); err != nil {
		return nil, err
	}

	userNull, err := useCase.userRepository.GetUserFromId(userId)
	if err != nil {
		return nil, err
	}

	var user models.User = *userNull

	var item models.Interval
	item.TaskId = taskId
	item.User = user.ToDTO()
	item.TimeStart = time.Now().Format(time.RFC3339)
	return useCase.intervalRepository.Add(item)

}

func (useCase *TimeIntervalUseCase) GetIntervalById(id, userId int) (*models.Interval, error) {
	interval, err := useCase.intervalRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	if err := useCase.taskUseCase.CheckUserHaveTask(interval.TaskId, userId); err != nil {
		return nil, err
	}

	return interval, err
}

func (useCase *TimeIntervalUseCase) GetIntervalsByTask(id, userId int) ([]models.Interval, error) {
	if err := useCase.taskUseCase.CheckUserHaveTask(id, userId); err != nil {
		return nil, err
	}

	return useCase.intervalRepository.GetByTaskId(id)

}

func (useCase *TimeIntervalUseCase) GetIntervalsByUser(userId int) ([]models.Interval, error) {
	return useCase.intervalRepository.GetByUserId(userId)
}

func (useCase *TimeIntervalUseCase) UpdateInterval(item models.Interval, userId int) error {
	if err := useCase.taskUseCase.CheckUserHaveTask(item.TaskId, userId); err != nil {
		return err
	}
	return useCase.intervalRepository.Update(item)
}

func (useCase *TimeIntervalUseCase) FinishInterval(taskId, userId int) error {
	if err := useCase.taskUseCase.CheckUserHaveTask(taskId, userId); err != nil {
		return err
	} else {
		item, err := useCase.intervalRepository.GetNotEndedInterval(taskId, userId)
		if err != nil {
			return err
		}
		item.TimeEnd = time.Now().Format(time.RFC3339)
		return useCase.intervalRepository.Update(*item)
	}

}
