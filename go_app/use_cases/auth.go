package usecases

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/zimnushka/task_me_go/go_app/app"
	"github.com/zimnushka/task_me_go/go_app/models"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

const salt = "79dbeb816582"

type AuthUseCase struct {
	userRepository repositories.UserRepository
}

func (useCase *AuthUseCase) CheckToken(token string) (*models.User, *app.AppError) {
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, app.NewError(http.StatusInternalServerError, app.ERR_Wrong_auth)
	}
	words := strings.Split(string(data), "-")
	if words[0] == salt {
		id, err := strconv.Atoi(words[len(words)-1])
		if err != nil {
			return nil, app.NewError(http.StatusInternalServerError, app.ERR_Wrong_auth)
		}
		data, err := useCase.userRepository.GetUserFromId(id)
		if err != nil {
			return nil, app.NewError(http.StatusNotFound, app.ERR_Wrong_auth)
		}
		return data, nil
	}
	return nil, app.NewError(http.StatusNotFound, app.ERR_Wrong_auth)
}

func (useCase *AuthUseCase) createToken(user models.User) string {
	id := *user.Id
	parametrs := fmt.Sprintf("%s-%d", salt, id)
	return base64.StdEncoding.EncodeToString([]byte(parametrs))

}

func (useCase *AuthUseCase) Register(user models.User) (string, *app.AppError) {
	if user.Email == "" || user.Password == "" {
		return "", app.NewError(http.StatusNotFound, app.ERR_Empty_field)
	}
	user.Id = nil
	userWithEmail, _ := useCase.userRepository.GetUserFromEmail(user.Email)
	if userWithEmail != nil {
		return "", app.NewError(http.StatusConflict, app.ERR_User_already_register)
	}
	newUser, err := useCase.userRepository.AddUser(user)
	return useCase.createToken(*newUser), app.AppErrorByError(err)
}
func (useCase *AuthUseCase) Login(email, password string) (string, *app.AppError) {
	if email == "" || password == "" {
		return "", app.NewError(http.StatusNotFound, app.ERR_Empty_field)
	}
	user, err := useCase.userRepository.GetUserFromEmail(email)

	if err != nil {
		return "", app.AppErrorByError(err)
	}
	if user.Password != password {
		return "", app.NewError(http.StatusNotFound, app.ERR_Not_found)
	}
	return useCase.createToken(*user), nil

}
