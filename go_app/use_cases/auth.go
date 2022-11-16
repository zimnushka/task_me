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

const salt = "79dbeb816582"

type AuthUseCase struct {
	userRepository repositories.UserRepository
}

func (useCase *AuthUseCase) CheckToken(token string) (*models.User, error) {
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}
	words := strings.Split(string(data), "-")
	if words[0] == salt {
		id, err := strconv.Atoi(words[len(words)-1])
		if err != nil {
			return nil, err
		}
		return useCase.userRepository.GetUserFromId(id)
	}
	return nil, errors.New("wrong auth")
}

func (useCase *AuthUseCase) createToken(user models.User) string {
	id := *user.Id
	parametrs := fmt.Sprintf("%s-%d", salt, id)
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
	newUser, err := useCase.userRepository.AddUser(user)
	return useCase.createToken(*newUser), err
}
func (useCase *AuthUseCase) Login(email, password string) (string, error) {
	if email == "" || password == "" {
		return "", errors.New("Password or email is empty")
	}
	user, err := useCase.userRepository.GetUserFromEmail(email)

	if err != nil {
		return "", err
	}
	if user.Password != password {
		return "", errors.New("User with this parametrs not found")
	}
	return useCase.createToken(*user), nil

}
