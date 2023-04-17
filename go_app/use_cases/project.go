package usecases

import (
	"net/http"

	"github.com/zimnushka/task_me_go/go_app/app_errors"
	"github.com/zimnushka/task_me_go/go_app/models"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

type ProjectUseCase struct {
	projectRepository     repositories.ProjectRepository
	projectUserRepository repositories.ProjectUserRepository
}

func (useCase *ProjectUseCase) GetProjectById(id, userId int) (*models.Project, *app_errors.AppError) {
	if err := useCase.CheckUserHaveProject(id, userId); err != nil {
		return nil, err
	}
	data, err := useCase.projectRepository.GetProjectFromId(id)
	if err != nil {
		return nil, app_errors.New(http.StatusNotFound, app_errors.ERR_Not_found)
	}
	return data, nil

}

func (useCase *ProjectUseCase) GetAllProjects(userId int) ([]models.Project, *app_errors.AppError) {
	data, err := useCase.projectUserRepository.GetProjectsByUser(userId)
	if err != nil {
		return nil, app_errors.New(http.StatusNotFound, app_errors.ERR_Not_found)
	}
	return data, nil
}

func (useCase *ProjectUseCase) GetProjectUsers(projectId, userId int) ([]models.User, *app_errors.AppError) {
	if err := useCase.CheckUserHaveProject(projectId, userId); err != nil {
		return nil, err
	}
	data, err := useCase.projectUserRepository.GetUsersByProject(projectId)

	if err != nil {
		return nil, app_errors.New(http.StatusNotFound, app_errors.ERR_Not_found)
	}
	return data, nil

}

func (useCase *ProjectUseCase) AddProject(project models.Project, userId int) (*models.Project, *app_errors.AppError) {
	if project.Title == "" {
		return nil, app_errors.New(http.StatusNotFound, app_errors.ERR_Empty_field)
	}
	project.Id = nil
	newProject, err := useCase.projectRepository.AddProject(project, userId)
	if err != nil {
		return nil, app_errors.FromError(err)
	}
	err = useCase.projectUserRepository.AddLink(*newProject.Id, userId)
	return newProject, app_errors.FromError(err)
}

func (useCase *ProjectUseCase) AddMemberToProject(projectId, userId, userRequestId int) *app_errors.AppError {
	if err := useCase.CheckUserHaveProject(projectId, userRequestId); err != nil {
		return err
	}
	return app_errors.FromError(useCase.projectUserRepository.AddLink(projectId, userId))

}

func (useCase *ProjectUseCase) UpdateProject(project models.Project, userId int) *app_errors.AppError {
	if err := useCase.CheckUserHaveProject(*project.Id, userId); err != nil {
		return err
	}
	return app_errors.FromError(useCase.projectRepository.UpdateProject(project))

}

func (useCase *ProjectUseCase) DeleteProject(id, userId int) *app_errors.AppError {
	if err := useCase.CheckUserHaveProject(id, userId); err != nil {
		return err
	}
	return app_errors.FromError(useCase.projectRepository.DeleteProject(id))

}

func (useCase *ProjectUseCase) DeleteMemberFromProject(projectId, userId, userRequestId int) *app_errors.AppError {
	if err := useCase.CheckUserHaveProject(projectId, userId); err != nil {
		return err
	}
	return app_errors.FromError(useCase.projectUserRepository.DeleteLink(projectId, userId))

}

func (useCase *ProjectUseCase) CheckUserHaveProject(projectId, userId int) *app_errors.AppError {
	projects, err := useCase.projectUserRepository.GetProjectsByUser(userId)
	if err != nil {
		return app_errors.New(http.StatusNotFound, app_errors.ERR_Not_found)
	}
	var id int
	for _, project := range projects {
		id = *project.Id
		if id == projectId {
			return nil
		}
	}
	return app_errors.New(http.StatusNotFound, app_errors.ERR_Forbiden)
}
