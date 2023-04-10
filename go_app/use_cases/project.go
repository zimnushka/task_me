package usecases

import (
	"errors"

	"github.com/zimnushka/task_me_go/go_app/app_errors"
	"github.com/zimnushka/task_me_go/go_app/models"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

type ProjectUseCase struct {
	projectRepository     repositories.ProjectRepository
	projectUserRepository repositories.ProjectUserRepository
}

func (useCase *ProjectUseCase) GetProjectById(id, userId int) (*models.Project, error) {
	if err := useCase.CheckUserHaveProject(id, userId); err != nil {
		return nil, err
	}
	return useCase.projectRepository.GetProjectFromId(id)

}

func (useCase *ProjectUseCase) GetAllProjects(userId int) ([]models.Project, error) {
	return useCase.projectUserRepository.GetProjectsByUser(userId)
}

func (useCase *ProjectUseCase) GetProjectUsers(projectId, userId int) ([]models.User, error) {
	if err := useCase.CheckUserHaveProject(projectId, userId); err != nil {
		return nil, err
	}
	return useCase.projectUserRepository.GetUsersByProject(projectId)

}

func (useCase *ProjectUseCase) AddProject(project models.Project, userId int) (*models.Project, error) {
	if project.Title == "" {
		return nil, errors.New(app_errors.ERR_Empty_field)
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
	if err := useCase.CheckUserHaveProject(projectId, userRequestId); err != nil {
		return err
	}
	return useCase.projectUserRepository.AddLink(projectId, userId)

}

func (useCase *ProjectUseCase) UpdateProject(project models.Project, userId int) error {
	if err := useCase.CheckUserHaveProject(*project.Id, userId); err != nil {
		return err
	}
	return useCase.projectRepository.UpdateProject(project)

}

func (useCase *ProjectUseCase) DeleteProject(id, userId int) error {
	if err := useCase.CheckUserHaveProject(id, userId); err != nil {
		return err
	}
	return useCase.projectRepository.DeleteProject(id)

}

func (useCase *ProjectUseCase) DeleteMemberFromProject(projectId, userId, userRequestId int) error {
	if err := useCase.CheckUserHaveProject(projectId, userId); err != nil {
		return err
	}
	return useCase.projectUserRepository.DeleteLink(projectId, userId)

}

func (useCase *ProjectUseCase) CheckUserHaveProject(projectId, userId int) error {
	projects, err := useCase.projectUserRepository.GetProjectsByUser(userId)
	if err != nil {
		return err
	}
	var id int
	for _, project := range projects {
		id = *project.Id
		if id == projectId {
			return nil
		}
	}
	return errors.New(app_errors.ERR_Forbiden)
}
