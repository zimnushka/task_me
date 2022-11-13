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
	models.Controller
}

func (controller UserController) Init() models.Controller {
	controller.Url = "/user/"
	controller.RegisterController("", controller.userHandler)
	return controller.Controller
}

func (controller UserController) userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		users, err := controller.userUseCase.GetAllUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		s, err := json.Marshal(users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, string(s))
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
