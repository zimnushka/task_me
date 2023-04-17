package usecases

import (
	"net/http"
	"time"

	"github.com/zimnushka/task_me_go/go_app/app"
	"github.com/zimnushka/task_me_go/go_app/models"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

type TaskUseCase struct {
	taskRepository     repositories.TaskRepository
	taskUserRepository repositories.TaskUserRepository

	projectUseCase ProjectUseCase
}

func (useCase *TaskUseCase) GetMembers(id, userId int) ([]models.User, *app.AppError) {
	if err := useCase.CheckUserHaveTask(id, userId); err != nil {
		return nil, err
	}
	data, err := useCase.taskUserRepository.GetUsersByTask(id)
	if err != nil {
		return nil, app.NewError(http.StatusNotFound, app.ERR_Not_found)
	}
	return data, nil
}

func (useCase *TaskUseCase) UpdateMembersList(id int, users []models.User, userId int) *app.AppError {
	if err := useCase.CheckUserHaveTask(id, userId); err != nil {
		return err
	}

	useCase.taskUserRepository.DeleteAllLinkByTask(id)
	for _, user := range users {
		if err := useCase.taskUserRepository.AddLink(id, *user.Id); err != nil {
			return app.AppErrorByError(err)
		}
	}
	return nil
}

func (useCase *TaskUseCase) GetTaskById(id, userId int) (*models.Task, *app.AppError) {
	task, err := useCase.taskRepository.GetTaskFromId(id)
	if err != nil {
		return nil, app.NewError(http.StatusNotFound, app.ERR_Not_found)
	}

	if err := useCase.projectUseCase.CheckUserHaveProject(task.ProjectId, userId); err != nil {
		return nil, err
	}

	task.Assigners = useCase.getAssignersIds(*task.Id)
	return task, nil

}

func (useCase *TaskUseCase) GetTaskByProjectId(projectId, userId int) ([]models.Task, *app.AppError) {
	if err := useCase.projectUseCase.CheckUserHaveProject(projectId, userId); err != nil {
		return nil, err
	}
	tasks, err := useCase.taskRepository.GetTasksFromProject(projectId)
	if err != nil {
		return nil, app.NewError(http.StatusNotFound, app.ERR_Not_found)
	}
	return useCase.addAssignersIdsToTaskList(tasks), nil

}

func (useCase *TaskUseCase) GetAllTasks(userId int) ([]models.Task, *app.AppError) {
	tasks, err := useCase.taskUserRepository.GetTasksByUser(userId)
	data, err := useCase.addAssignersIdsToTaskList(tasks), err
	if err != nil {
		return nil, app.NewError(http.StatusNotFound, app.ERR_Not_found)
	}
	return data, nil

}

func (useCase *TaskUseCase) AddTask(task models.Task, userId int) (*models.Task, *app.AppError) {
	if task.Title == "" {
		return nil, app.NewError(http.StatusNotFound, app.ERR_Empty_field)
	}
	task.Id = nil
	task.StartDate = time.Now().Format(time.RFC3339)
	data, err := useCase.taskRepository.AddTask(task)
	if err != nil {
		return nil, app.AppErrorByError(err)
	}
	return data, nil
}

func (useCase *TaskUseCase) UpdateTask(item models.Task, userId int) *app.AppError {
	if err := useCase.CheckUserHaveTask(*item.Id, userId); err != nil {
		return err
	}
	return app.AppErrorByError(useCase.taskRepository.UpdateTask(item))

}

func (useCase *TaskUseCase) DeleteTask(id, userId int) *app.AppError {
	if err := useCase.CheckUserHaveTask(id, userId); err != nil {
		return err
	}
	return app.AppErrorByError(useCase.taskRepository.DeleteTask(id))

}

func (useCase *TaskUseCase) addAssignersIdsToTaskList(tasks []models.Task) []models.Task {
	var newList []models.Task
	for _, task := range tasks {
		usersIds := useCase.getAssignersIds(*task.Id)
		task.Description = ""
		task.Assigners = usersIds
		newList = append(newList, task)

	}
	return newList
}

func (useCase *TaskUseCase) getAssignersIds(taskId int) []int {
	users, err := useCase.taskUserRepository.GetUsersByTask(taskId)
	if err == nil {
		var usersIds []int
		for _, element := range users {
			usersIds = append(usersIds, *element.Id)
		}
		return usersIds
	}
	var empty []int
	return empty
}

func (useCase *TaskUseCase) CheckUserHaveTask(taskId, userId int) *app.AppError {
	item, _ := useCase.taskRepository.GetTaskFromId(taskId)
	projects, err := useCase.projectUseCase.GetAllProjects(userId)
	if err != nil || item == nil {
		return err
	}
	var id int
	for _, project := range projects {
		id = *project.Id
		if id == item.ProjectId {
			return nil
		}
	}
	return app.NewError(http.StatusForbidden, app.ERR_Forbiden)
}
