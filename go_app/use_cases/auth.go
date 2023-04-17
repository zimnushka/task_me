package usecases

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/zimnushka/task_me_go/go_app/app_errors"
	"github.com/zimnushka/task_me_go/go_app/models"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

const salt = "79dbeb816582"

type AuthUseCase struct {
	userRepository repositories.UserRepository
}

func (useCase *AuthUseCase) CheckToken(token string) (*models.User, *app_errors.AppError) {
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, app_errors.New(http.StatusInternalServerError, app_errors.ERR_Wrong_auth)
	}
	words := strings.Split(string(data), "-")
	if words[0] == salt {
		id, err := strconv.Atoi(words[len(words)-1])
		if err != nil {
			return nil, app_errors.New(http.StatusInternalServerError, app_errors.ERR_Wrong_auth)
		}
		data, err := useCase.userRepository.GetUserFromId(id)
		if err != nil {
			return nil, app_errors.New(http.StatusNotFound, app_errors.ERR_Wrong_auth)
		}
		return data, nil
	}
	return nil, app_errors.New(http.StatusNotFound, app_errors.ERR_Wrong_auth)
}

func (useCase *AuthUseCase) createToken(user models.User) string {
	id := *user.Id
	parametrs := fmt.Sprintf("%s-%d", salt, id)
	return base64.StdEncoding.EncodeToString([]byte(parametrs))

}

func (useCase *AuthUseCase) Register(user models.User) (string, *app_errors.AppError) {
	if user.Email == "" || user.Password == "" {
		return "", app_errors.New(http.StatusNotFound, app_errors.ERR_Empty_field)
	}
	user.Id = nil
	userWithEmail, _ := useCase.userRepository.GetUserFromEmail(user.Email)
	if userWithEmail != nil {
		return "", app_errors.New(http.StatusConflict, app_errors.ERR_User_already_register)
	}
	newUser, err := useCase.userRepository.AddUser(user)
	return useCase.createToken(*newUser), app_errors.FromError(err)
}
func (useCase *AuthUseCase) Login(email, password string) (string, *app_errors.AppError) {
	if email == "" || password == "" {
		return "", app_errors.New(http.StatusNotFound, app_errors.ERR_Empty_field)
	}
	user, err := useCase.userRepository.GetUserFromEmail(email)

	if err != nil {
		return "", app_errors.FromError(err)
	}
	if user.Password != password {
		return "", app_errors.New(http.StatusNotFound, app_errors.ERR_Not_found)
	}
	return useCase.createToken(*user), nil

}
