package usecases

import (
	"net/http"
	"time"

	"github.com/zimnushka/task_me_go/go_app/app"
	"github.com/zimnushka/task_me_go/go_app/models"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

type TimeIntervalUseCase struct {
	intervalRepository repositories.IntervalRepository
	taskRepository     repositories.TaskRepository

	taskUseCase    TaskUseCase
	projectUseCase ProjectUseCase
}

func (useCase *TimeIntervalUseCase) AddInterval(taskId int, user models.User) (*models.Interval, *app.AppError) {
	notEndedInterval, _ := useCase.intervalRepository.GetNotEndedInterval(*user.Id)
	if notEndedInterval != nil {
		return nil, app.NewError(http.StatusInternalServerError, app.ERR_TimeInterval_Already_Started)
	}
	taskData, err := useCase.taskRepository.GetTaskFromId(taskId)
	if err != nil {
		return nil, app.NewError(http.StatusNotFound, app.ERR_Not_found)
	}
	task := *taskData

	if err := useCase.taskUseCase.CheckUserHaveTaskById(taskId, *user.Id); err != nil {
		return nil, err
	}

	var item models.Interval
	item.Task = task.ToDTO()
	item.User = user.ToDTO()
	item.TimeStart = time.Now().Format(time.RFC3339)
	data, err := useCase.intervalRepository.Add(item)
	if err != nil {
		return nil, app.AppErrorByError(err)
	}
	return data, nil
}

func (useCase *TimeIntervalUseCase) GetIntervalById(id, userId int) (*models.Interval, *app.AppError) {
	interval, err := useCase.intervalRepository.GetById(id)
	if err != nil {
		return nil, app.NewError(http.StatusNotFound, app.ERR_Not_found)
	}
	if err := useCase.taskUseCase.CheckUserHaveTaskById(*interval.Task.Id, userId); err != nil {
		return nil, err
	}

	return interval, app.NewError(http.StatusNotFound, app.ERR_Not_found)
}

func (useCase *TimeIntervalUseCase) GetIntervalsByTask(id, userId int) ([]models.Interval, *app.AppError) {
	if err := useCase.taskUseCase.CheckUserHaveTaskById(id, userId); err != nil {
		return nil, err
	}
	data, err := useCase.intervalRepository.GetByTaskId(id)
	if err != nil {
		return nil, app.NewError(http.StatusNotFound, app.ERR_Not_found)
	}
	return data, nil
}

func (useCase *TimeIntervalUseCase) GetIntervalsByProject(id, userId int) ([]models.Interval, *app.AppError) {
	if err := useCase.projectUseCase.CheckUserHaveProject(id, userId); err != nil {
		return nil, err
	}
	data, err := useCase.intervalRepository.GetByProjectId(id)
	if err != nil {
		return nil, app.NewError(http.StatusNotFound, app.ERR_Not_found)
	}
	return data, nil
}

func (useCase *TimeIntervalUseCase) GetIntervalsByUser(userId int) ([]models.Interval, *app.AppError) {
	data, err := useCase.intervalRepository.GetByUserId(userId)
	if err != nil {
		return nil, app.NewError(http.StatusNotFound, app.ERR_Not_found)
	}
	return data, nil
}

func (useCase *TimeIntervalUseCase) UpdateInterval(item models.Interval, userId int) *app.AppError {
	if err := useCase.taskUseCase.CheckUserHaveTaskById(*item.Task.Id, userId); err != nil {
		return err
	}
	return app.AppErrorByError(useCase.intervalRepository.Update(item))
}

func (useCase *TimeIntervalUseCase) FinishInterval(userId int, desc string) *app.AppError {
	item, err := useCase.intervalRepository.GetNotEndedInterval(userId)
	if err != nil {
		return app.NewError(http.StatusNotFound, app.ERR_Not_found)
	}
	item.TimeEnd = time.Now().Format(time.RFC3339)
	item.Description = desc
	return app.AppErrorByError(useCase.intervalRepository.Update(*item))
}
