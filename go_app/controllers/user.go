package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/zimnushka/task_me_go/go_app/models"
	usecases "github.com/zimnushka/task_me_go/go_app/use_cases"
)

type UserController struct {
	userUseCase usecases.UserUseCase
	authUseCase usecases.AuthUseCase
	corsUseCase usecases.CorsUseCase
	models.Controller
}

func (controller UserController) Init() models.Controller {
	controller.Url = "/user/"
	controller.RegisterController("", controller.userHandler)
	return controller.Controller
}

func (controller UserController) userHandler(w http.ResponseWriter, r *http.Request) {
	controller.corsUseCase.DisableCors(&w, r)

	user, err := controller.authUseCase.CheckToken(r.Header.Get(models.HeaderAuth))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		var jsonData []byte
		path := strings.TrimPrefix(r.URL.Path, controller.Url)
		if path == "me" {
			jsonData, err = json.Marshal(user)
		} else {
			id, err := strconv.Atoi(path)
			if err == nil {
				user, err := controller.userUseCase.GetUserById(id)
				if err != nil {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				jsonData, err = json.Marshal(user)

			} else {
				err = nil
				users, err := controller.userUseCase.GetAllUsers()
				if err != nil {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				jsonData, err = json.Marshal(users)

			}
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonData))
	case "POST":
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		newUser, err := controller.userUseCase.AddUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		jsonData, err := json.Marshal(newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonData))
	case "PUT":
		var user models.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		_, err = controller.userUseCase.UpdateUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	case "DELETE":
		idString := strings.TrimPrefix(r.URL.Path, controller.Url)
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = controller.userUseCase.DeleteUser(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	}

}
