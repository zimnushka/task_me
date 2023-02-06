package usecases

import (
	"errors"

	"github.com/zimnushka/task_me_go/go_app/models"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

type TaskUseCase struct {
	taskRepository     repositories.TaskRepository
	taskUserRepository repositories.TaskUserRepository

	projectUseCase ProjectUseCase
}

func (useCase *TaskUseCase) AddMember(id, newMember, userId int) error {
	access, err := useCase.CheckUserHaveTask(id, userId)
	if err != nil {
		return err
	}
	if access {
		return useCase.taskUserRepository.AddLink(id, newMember)
	}
	return errors.New("Forbiden")
}

func (useCase *TaskUseCase) GetMembers(id, userId int) ([]models.User, error) {
	access, err := useCase.CheckUserHaveTask(id, userId)
	if err != nil {
		return nil, err
	}
	if access {
		return useCase.taskUserRepository.GetUsersByTask(id)
	}
	return nil, errors.New("Forbiden")
}

func (useCase *TaskUseCase) DeleteMember(id, memberId, userId int) error {
	access, err := useCase.CheckUserHaveTask(id, userId)
	if err != nil {
		return err
	}
	if access {
		return useCase.taskUserRepository.DeleteLink(id, memberId)
	}
	return errors.New("Forbiden")
}

func (useCase *TaskUseCase) GetTaskById(id, userId int) (*models.Task, error) {
	task, err := useCase.taskRepository.GetTaskFromId(id)
	if err != nil {
		return nil, err
	}

	access, err := useCase.projectUseCase.CheckUserHaveProject(task.ProjectId, userId)
	if err != nil {
		return nil, err
	}
	if access {
		return task, err
	}
	return nil, errors.New("Forbiden")
}

func (useCase *TaskUseCase) GetTaskByProjectId(projectId, userId int) ([]models.Task, error) {
	access, err := useCase.projectUseCase.CheckUserHaveProject(projectId, userId)
	if err != nil {
		return nil, err
	}
	if access {
		return useCase.taskRepository.GetTasksFromProject(projectId)
	}
	return nil, errors.New("Forbiden")
}

func (useCase *TaskUseCase) GetAllTasks(userId int) ([]models.Task, error) {
	return useCase.taskUserRepository.GetTasksByUser(userId)
}

func (useCase *TaskUseCase) AddTask(task models.Task, userId int) (*models.Task, error) {
	if task.Title == "" {
		return nil, errors.New("Title is empty")
	}
	task.Id = nil
	return useCase.taskRepository.AddTask(task)
}

func (useCase *TaskUseCase) UpdateTask(item models.Task, userId int) error {
	access, err := useCase.CheckUserHaveTask(*item.Id, userId)
	if err != nil {
		return err
	}
	if access {
		return useCase.taskRepository.UpdateTask(item)
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

func (useCase *TaskUseCase) CheckUserHaveTask(taskId, userId int) (bool, error) {
	item, err := useCase.taskRepository.GetTaskFromId(taskId)
	projects, err := useCase.projectUseCase.GetAllProjects(userId)
	if err != nil || item == nil {
		return false, err
	}
	var id int
	for _, project := range projects {
		id = *project.Id
		if id == *&item.ProjectId {
			return true, nil
		}
	}
	return false, nil
}
