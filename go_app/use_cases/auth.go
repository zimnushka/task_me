package usecases

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/zimnushka/task_me_go/go_app/models"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

const salt = "79dbeb81-6582-419d-a61a-ba5a75a293c7"

type AuthUseCase struct {
	userRepository repositories.UserRepository
}

func (useCase *AuthUseCase) CheckToken(token string) (*models.User, error) {
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}
	words := strings.Split(fmt.Sprintf("%q\n", data), "//")
	if words[len(words)-1] == salt {
		id, err := strconv.Atoi(words[0])
		if err != nil {
			return nil, err
		}
		return useCase.userRepository.GetUserFromId(id)
	}
	return nil, errors.New("wrong auth")
}

func (useCase *AuthUseCase) createToken(user models.User) string {

	parametrs := fmt.Sprintf("%s//%d", salt, user.Id)
	return base64.StdEncoding.EncodeToString([]byte(parametrs))

}

func (useCase *AuthUseCase) Register(user models.User) (string, error) {
	if user.Email == "" || user.Password == "" {
		return "", errors.New("Password or email is empty")
	}
	user.Id = nil
	userWithEmail, _ := useCase.userRepository.GetUserFromEmail(user.Email)
	if userWithEmail != nil {
		return "", errors.New("User with this email was created")
	}
	err := useCase.userRepository.AddUser(user)
	if err != nil {
		return "", err
	}
	newUser, err := useCase.userRepository.GetUserFromEmail(user.Email)

	return useCase.createToken(*newUser), nil
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
