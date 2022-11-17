package usecases

import (
	"errors"

	"github.com/zimnushka/task_me_go/go_app/models"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

type TaskUseCase struct {
	taskRepository     repositories.TaskRepository
	taskUserRepository repositories.TaskUserRepository
}

func (useCase *TaskUseCase) GetTaskById(id, userId int) (*models.Task, error) {
	access, err := useCase.CheckUserHaveTask(id, userId)
	if err != nil {
		return nil, err
	}
	if access {
		return useCase.taskRepository.GetTaskFromId(id)
	}
	return nil, errors.New("Forbiden")
}

func (useCase *TaskUseCase) GetAllTasks(userId int) ([]models.Task, error) {
	return useCase.taskUserRepository.GetTasksByUser(userId)
}

func (useCase *TaskUseCase) GetTaskUsers(projectId, userId int) ([]models.User, error) {
	access, err := useCase.CheckUserHaveTask(projectId, userId)
	if err != nil {
		return nil, err
	}
	if access {
		return useCase.taskUserRepository.GetUsersByTask(projectId)
	}
	return nil, errors.New("Forbiden")
}

func (useCase *TaskUseCase) AddTask(project models.Task, userId int) (*models.Task, error) {
	if project.Title == "" {
		return nil, errors.New("Title is empty")
	}
	project.Id = nil
	newTask, err := useCase.taskRepository.AddTask(project)
	if err != nil {
		return nil, err
	}
	err = useCase.taskUserRepository.AddLink(*newTask.Id, userId)
	return newTask, err
}

func (useCase *TaskUseCase) AddMemberToTask(projectId, userId, userRequestId int) error {
	access, err := useCase.CheckUserHaveTask(projectId, userRequestId)
	if err != nil {
		return err
	}
	if access {
		return useCase.taskUserRepository.AddLink(projectId, userId)
	}
	return errors.New("Forbiden")
}

func (useCase *TaskUseCase) UpdateTask(project models.Task, userId int) error {
	access, err := useCase.CheckUserHaveTask(*project.Id, userId)
	if err != nil {
		return err
	}
	if access {
		return useCase.taskRepository.UpdateTask(project)
	}
	return errors.New("Forbiden")
}

func (useCase *TaskUseCase) DeleteTask(id, userId int) error {
	access, err := useCase.CheckUserHaveTask(id, userId)
	if err != nil {
		return err
	}
	if access {
		return useCase.taskRepository.DeleteTask(id)
	}
	return errors.New("Forbiden")
}

func (useCase *TaskUseCase) DeleteMemberFromTask(projectId, userId, userRequestId int) error {
	access, err := useCase.CheckUserHaveTask(projectId, userId)
	if err != nil {
		return err
	}
	if access {
		return useCase.taskUserRepository.DeleteLink(projectId, userId)
	}
	return errors.New("Forbiden")
}

func (useCase *TaskUseCase) CheckUserHaveTask(projectId, userId int) (bool, error) {
	projects, err := useCase.taskUserRepository.GetTasksByUser(userId)
	if err != nil {
		return false, err
	}
	var id int
	for _, project := range projects {
		id = *project.Id
		if id == projectId {
			return true, nil
		}
	}
	return false, nil
}
