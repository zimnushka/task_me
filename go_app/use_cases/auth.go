package usecases

import (
	"errors"

	"github.com/zimnushka/task_me_go/go_app/models"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

type AuthUseCase struct {
	userRepository repositories.UserRepository
}

func (useCase *AuthUseCase) Register(user models.User) (*models.User, error) {
	if user.Email == "" || user.Password == "" {
		return nil, errors.New("Password or email is empty")
	}
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
func (useCase *AuthUseCase) Login(email, password string) (*models.User, error) {
	if email == "" || password == "" {
		return nil, errors.New("Password or email is empty")
	}
	user, err := useCase.userRepository.GetUserFromEmail(email)

	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, errors.New("User with this parametrs not found")
	}
	return user, nil

}
