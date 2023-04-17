package usecases

import (
	"net/http"

	"github.com/zimnushka/task_me_go/go_app/app"
	"github.com/zimnushka/task_me_go/go_app/models"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

type UserUseCase struct {
	userRepository repositories.UserRepository
}

func (useCase *UserUseCase) GetUserById(id int) (*models.User, *app.AppError) {
	data, err := useCase.userRepository.GetUserFromId(id)
	if err != nil {
		return nil, app.NewError(http.StatusNotFound, app.ERR_Not_found)
	}
	return data, nil
}

func (useCase *UserUseCase) GetUserByEmail(email string) (*models.User, *app.AppError) {
	data, err := useCase.userRepository.GetUserFromEmail(email)
	if err != nil {
		return nil, app.NewError(http.StatusNotFound, app.ERR_Not_found)
	}
	return data, nil
}

func (useCase *UserUseCase) GetAllUsers() ([]models.User, *app.AppError) {
	data, err := useCase.userRepository.GetUsers()
	if err != nil {
		return nil, app.NewError(http.StatusNotFound, app.ERR_Not_found)
	}
	return data, nil
}
func (useCase *UserUseCase) AddUser(user models.User) (*models.User, *app.AppError) {
	user.Id = nil
	userWithEmail, _ := useCase.userRepository.GetUserFromEmail(user.Email)
	if userWithEmail != nil {
		return nil, app.NewError(http.StatusNotFound, app.ERR_Not_found)
	}

	data, err := useCase.userRepository.AddUser(user)
	if err != nil {
		return nil, app.AppErrorByError(err)
	}
	return data, nil
}
func (useCase *UserUseCase) UpdateUser(user models.User, userId int) (*models.User, *app.AppError) {
	appErr := useCase.CheckUserHaveAcces(*user.Id, userId)
	if appErr != nil {
		return nil, appErr
	}
	if user.Password == "" {
		userWithPass, _ := useCase.userRepository.GetUserFromId(*user.Id)
		user.Password = userWithPass.Password
	}

	err := useCase.userRepository.UpdateUser(user)
	if err != nil {
		return nil, app.AppErrorByError(err)
	}
	data, err := useCase.userRepository.GetUserFromId(*user.Id)
	if err != nil {
		return nil, app.AppErrorByError(err)
	}
	return data, nil
}
func (useCase *UserUseCase) DeleteUser(id, userId int) *app.AppError {
	err := useCase.CheckUserHaveAcces(id, userId)
	if err != nil {
		return err
	}
	return app.AppErrorByError(useCase.userRepository.DeleteUser(id))
}

func (useCase *UserUseCase) CheckUserHaveAcces(userEditedId, userId int) *app.AppError {
	if userEditedId == userId {
		return nil
	}
	return app.NewError(http.StatusForbidden, app.ERR_Forbiden)
}
