package usecases

import (
	"encoding/json"
	"fmt"

	"github.com/zimnushka/task_me_go/go_app/models"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

type AuthUseCase struct{}

func (useCase *AuthUseCase) Register(user models.User) (string, error) {
	user.Id = nil
	repositories.AddUser(user)
	query := fmt.Sprintf("SELECT * FROM users WHERE email = '%s'", user.Email)
	user = *repositories.GetUser(query)
	s, e := json.Marshal(user)
	return string(s), e
}
func (useCase *AuthUseCase) Login(email, password string) (string, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE email = '%s' AND password = '%s'", email, password)
	user := repositories.GetUser(query)
	if user == nil {
		var err error
		return "", err
	}
	s, e := json.Marshal(user)
	return string(s), e

}
