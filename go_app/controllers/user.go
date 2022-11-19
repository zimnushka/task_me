package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/zimnushka/task_me_go/go_app/models"
	usecases "github.com/zimnushka/task_me_go/go_app/use_cases"
)

type UserController struct {
	userUseCase usecases.UserUseCase
	authUseCase usecases.AuthUseCase
	models.Controller
}

func (controller UserController) Init() models.Controller {
	controller.Url = "/user/"
	controller.RegisterController("", controller.userHandler)
	return controller.Controller
}

func (controller UserController) userHandler(w http.ResponseWriter, r *http.Request) {
	user, err := controller.authUseCase.CheckToken(r.Header.Get(models.HeaderAuth))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	switch r.Method {
	case "GET":
		var jsonData []byte
		path := strings.TrimPrefix(r.URL.Path, controller.Url)
		if path == "me" {
			user, err := controller.userUseCase.GetUserById(*user.Id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			jsonData, err = json.Marshal(user)
			return
		}
		id, err := strconv.Atoi(path)
		if err == nil {
			user, err := controller.userUseCase.GetUserById(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			jsonData, err = json.Marshal(user)

		} else {
			err = nil
			users, err := controller.userUseCase.GetAllUsers()
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			jsonData, err = json.Marshal(users)

		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, string(jsonData))
	case "POST":
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newUser, err := controller.userUseCase.AddUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		s, err := json.Marshal(newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, string(s))
	case "PUT":
		var user models.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = controller.userUseCase.UpdateUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "")
	case "DELETE":
		idString := strings.TrimPrefix(r.URL.Path, controller.Url)
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = controller.userUseCase.DeleteUser(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "")
	}

}
