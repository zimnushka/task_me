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
func (useCase *ProjectUseCase) AddProject(project models.Project, userId int) (*models.Project, error) {
	if project.Title == "" {
		return nil, errors.New("Title is empty")
	}
	project.Id = nil
	newProject, err := useCase.projectRepository.AddProject(project)
	if err != nil {
		return nil, err
	}
	err = useCase.projectUserRepository.AddLink(*newProject.Id, userId)
	return newProject, err

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

func (useCase *ProjectUseCase) CheckUserHaveProject(projectId, userId int) (bool, error) {
	projects, err := useCase.projectUserRepository.GetProjectsByUser(userId)
	if err != nil {
		return false, err
	}
	for _, project := range projects {
		if project.Id == &projectId {
			return true, nil
		}
	}
	return false, nil
}
