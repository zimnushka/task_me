package usecases

import (
	"net/http"
	"time"

	"github.com/zimnushka/task_me_go/go_app/app_errors"
	"github.com/zimnushka/task_me_go/go_app/models"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

type TimeIntervalUseCase struct {
	intervalRepository repositories.IntervalRepository
	userRepository     repositories.UserRepository

	taskUseCase TaskUseCase
}

func (useCase *TimeIntervalUseCase) AddInterval(taskId, userId int) (*models.Interval, *app_errors.AppError) {
	if err := useCase.taskUseCase.CheckUserHaveTask(taskId, userId); err != nil {
		return nil, err
	}

	userNull, err := useCase.userRepository.GetUserFromId(userId)
	if err != nil {
		return nil, app_errors.New(http.StatusNotFound, app_errors.ERR_Not_found)
	}

	var user models.User = *userNull

	var item models.Interval
	item.TaskId = taskId
	item.User = user.ToDTO()
	item.TimeStart = time.Now().Format(time.RFC3339)
	data, err := useCase.intervalRepository.Add(item)
	if err != nil {
		return nil, app_errors.FromError(err)
	}
	return data, nil

}

func (useCase *TimeIntervalUseCase) GetIntervalById(id, userId int) (*models.Interval, *app_errors.AppError) {
	interval, err := useCase.intervalRepository.GetById(id)
	if err != nil {
		return nil, app_errors.New(http.StatusNotFound, app_errors.ERR_Not_found)
	}
	if err := useCase.taskUseCase.CheckUserHaveTask(interval.TaskId, userId); err != nil {
		return nil, err
	}

	return interval, app_errors.New(http.StatusNotFound, app_errors.ERR_Not_found)
}

func (useCase *TimeIntervalUseCase) GetIntervalsByTask(id, userId int) ([]models.Interval, *app_errors.AppError) {
	if err := useCase.taskUseCase.CheckUserHaveTask(id, userId); err != nil {
		return nil, err
	}
	data, err := useCase.intervalRepository.GetByTaskId(id)
	if err != nil {
		return nil, app_errors.New(http.StatusNotFound, app_errors.ERR_Not_found)
	}
	return data, nil
}

func (useCase *TimeIntervalUseCase) GetIntervalsByUser(userId int) ([]models.Interval, *app_errors.AppError) {
	data, err := useCase.intervalRepository.GetByUserId(userId)
	if err != nil {
		return nil, app_errors.New(http.StatusNotFound, app_errors.ERR_Not_found)
	}
	return data, nil
}

func (useCase *TimeIntervalUseCase) UpdateInterval(item models.Interval, userId int) *app_errors.AppError {
	if err := useCase.taskUseCase.CheckUserHaveTask(item.TaskId, userId); err != nil {
		return err
	}
	return app_errors.FromError(useCase.intervalRepository.Update(item))
}

func (useCase *TimeIntervalUseCase) FinishInterval(taskId, userId int) *app_errors.AppError {
	if err := useCase.taskUseCase.CheckUserHaveTask(taskId, userId); err != nil {
		return err
	} else {
		item, err := useCase.intervalRepository.GetNotEndedInterval(taskId, userId)
		if err != nil {
			return app_errors.New(http.StatusNotFound, app_errors.ERR_Not_found)
		}
		item.TimeEnd = time.Now().Format(time.RFC3339)
		return app_errors.FromError(useCase.intervalRepository.Update(*item))
	}

}
