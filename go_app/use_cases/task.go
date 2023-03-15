package usecases

import (
	"errors"
	"time"

	"github.com/zimnushka/task_me_go/go_app/models"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

type TaskUseCase struct {
	taskRepository     repositories.TaskRepository
	taskUserRepository repositories.TaskUserRepository

	projectUseCase ProjectUseCase
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

func (useCase *TaskUseCase) UpdateMembersList(id int, users []models.User, userId int) error {
	access, err := useCase.CheckUserHaveTask(id, userId)
	if err != nil {
		return err
	}
	if access {
		useCase.taskUserRepository.DeleteAllLinkByTask(id)
		for _, user := range users {
			err = useCase.taskUserRepository.AddLink(id, *user.Id)
			if err != nil {
				return err
			}
		}

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
		task.Assigners = useCase.getAssignersIds(*task.Id)
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
		tasks, err := useCase.taskRepository.GetTasksFromProject(projectId)
		return useCase.addAssignersIdsToTaskList(tasks), err
	}
	return nil, errors.New("Forbiden")
}

func (useCase *TaskUseCase) GetAllTasks(userId int) ([]models.Task, error) {
	tasks, err := useCase.taskUserRepository.GetTasksByUser(userId)
	return useCase.addAssignersIdsToTaskList(tasks), err
}

func (useCase *TaskUseCase) AddTask(task models.Task, userId int) (*models.Task, error) {
	if task.Title == "" {
		return nil, errors.New("Title is empty")
	}
	task.Id = nil
	task.StartDate = time.Now().Format(time.RFC3339)
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
