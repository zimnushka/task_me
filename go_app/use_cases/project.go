package usecases

import (
	"errors"

	"github.com/zimnushka/task_me_go/go_app/models"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

type ProjectUseCase struct {
	projectRepository     repositories.ProjectRepository
	projectUserRepository repositories.ProjectUserRepository
}

func (useCase *ProjectUseCase) GetProjectById(id, userId int) (*models.Project, error) {
	access, err := useCase.CheckUserHaveProject(id, userId)
	if err != nil {
		return nil, err
	}
	if access {
		return useCase.projectRepository.GetProjectFromId(id)
	}
	return nil, errors.New("Forbiden")
}

func (useCase *ProjectUseCase) GetAllProjects(userId int) ([]models.Project, error) {
	return useCase.projectUserRepository.GetProjectsByUser(userId)
}

func (useCase *ProjectUseCase) GetProjectUsers(projectId, userId int) ([]models.User, error) {
	access, err := useCase.CheckUserHaveProject(projectId, userId)
	if err != nil {
		return nil, err
	}
	if access {
		return useCase.projectUserRepository.GetUsersByProject(projectId)
	}
	return nil, errors.New("Forbiden")
}

func (useCase *ProjectUseCase) AddProject(project models.Project, userId int) (*models.Project, error) {
	if project.Title == "" {
		return nil, errors.New("Title is empty")
	}
	project.Id = nil
	newProject, err := useCase.projectRepository.AddProject(project, userId)
	if err != nil {
		return nil, err
	}
	err = useCase.projectUserRepository.AddLink(*newProject.Id, userId)
	return newProject, err
}

func (useCase *ProjectUseCase) AddMemberToProject(projectId, userId, userRequestId int) error {
	access, err := useCase.CheckUserHaveProject(projectId, userRequestId)
	if err != nil {
		return err
	}
	if access {
		return useCase.projectUserRepository.AddLink(projectId, userId)
	}
	return errors.New("Forbiden")
}

func (useCase *ProjectUseCase) UpdateProject(project models.Project, userId int) error {
	access, err := useCase.CheckUserHaveProject(*project.Id, userId)
	if err != nil {
		return err
	}
	if access {
		return useCase.projectRepository.UpdateProject(project)
	}
	return errors.New("Forbiden")
}

func (useCase *ProjectUseCase) DeleteProject(id, userId int) error {
	access, err := useCase.CheckUserHaveProject(id, userId)
	if err != nil {
		return err
	}
	if access {
		return useCase.projectRepository.DeleteProject(id)
	}
	return errors.New("Forbiden")
}

func (useCase *ProjectUseCase) DeleteMemberFromProject(projectId, userId, userRequestId int) error {
	access, err := useCase.CheckUserHaveProject(projectId, userId)
	if err != nil {
		return err
	}
	if access {
		return useCase.projectUserRepository.DeleteLink(projectId, userId)
	}
	return errors.New("Forbiden")
}

func (useCase *ProjectUseCase) CheckUserHaveProject(projectId, userId int) (bool, error) {
	projects, err := useCase.projectUserRepository.GetProjectsByUser(userId)
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
