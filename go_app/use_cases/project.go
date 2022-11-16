package usecases

import (
	"github.com/zimnushka/task_me_go/go_app/models"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

type ProjectUseCase struct {
	projectRepository     repositories.ProjectRepository
	projectUserRepository repositories.ProjectUserRepository
}

func (useCase *ProjectUseCase) GetProjectById(id, userId int) (*models.Project, error) {
	return useCase.projectRepository.GetProjectFromId(id)
}

func (useCase *ProjectUseCase) GetAllProjects(userId int) ([]models.Project, error) {
	return useCase.projectRepository.GetProjects()
}
func (useCase *ProjectUseCase) AddProject(project models.Project, userId int) (*models.Project, error) {
	project.Id = nil

	return useCase.projectRepository.AddProject(project)

}
func (useCase *ProjectUseCase) UpdateProject(project models.Project, userId int) (*models.Project, error) {
	err := useCase.projectRepository.UpdateProject(project)
	if err != nil {
		return nil, err
	}
	return useCase.projectRepository.GetProjectFromId(*project.Id)
}

func (useCase *ProjectUseCase) DeleteProject(id, userId int) error {
	return useCase.projectRepository.DeleteProject(id)
}
