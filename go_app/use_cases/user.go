package usecases

import (
	"errors"

	"github.com/zimnushka/task_me_go/go_app/models"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

type UserUseCase struct {
	userRepository repositories.UserRepository
}

func (useCase *UserUseCase) GetUserById(id int) (*models.User, error) {
	return useCase.userRepository.GetUserFromId(id)
}

func (useCase *UserUseCase) GetAllUsers() ([]models.User, error) {
	return useCase.userRepository.GetUsers()
}
func (useCase *UserUseCase) AddUser(user models.User) (*models.User, error) {
	user.Id = nil
	userWithEmail, _ := useCase.userRepository.GetUserFromEmail(user.Email)
	if userWithEmail != nil {
		return nil, errors.New("User with this email was created")
	}
	err := useCase.userRepository.AddUser(user)
	if err != nil {
		return nil, err
	}
	return useCase.userRepository.GetUserFromEmail(user.Email)
}
func (useCase *UserUseCase) UpdateUser(user models.User) (*models.User, error) {
	err := useCase.userRepository.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return useCase.userRepository.GetUserFromId(*user.Id)
}
func (useCase *UserUseCase) DeleteUser(id int) error {
	return useCase.userRepository.DeleteUser(id)
}
