package controllers

import (
	"encoding/json"

	"net/http"
	"strconv"
	"strings"

	"github.com/zimnushka/task_me_go/go_app/models"
	usecases "github.com/zimnushka/task_me_go/go_app/use_cases"
)

type TaskMemberController struct {
	authUseCase usecases.AuthUseCase
	taskUseCase usecases.TaskUseCase
	corsUseCase usecases.CorsUseCase
	models.Controller
}

func (controller TaskMemberController) Init() models.Controller {
	controller.Url = "/taskMembers/"
	controller.RegisterController("", controller.taskHandler)
	return controller.Controller
}

func (controller TaskMemberController) taskHandler(w http.ResponseWriter, r *http.Request) {
	controller.corsUseCase.DisableCors(&w, r)

	user, err := controller.authUseCase.CheckToken(r.Header.Get(models.HeaderAuth))
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	idString := strings.Split(r.URL.Path, "/")
	taskId, err := strconv.Atoi(idString[len(idString)-1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		users, err := controller.taskUseCase.GetMembers(taskId, *user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		jsonData, err := json.Marshal(users)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonData))
	case "PUT":
		var user_id int
		user_id_string := r.URL.Query().Get("user_id")
		user_id, err = strconv.Atoi(user_id_string)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		err := controller.taskUseCase.AddMember(taskId, user_id, *user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	case "DELETE":
		var user_id int
		user_id_string := r.URL.Query().Get("user_id")
		user_id, err = strconv.Atoi(user_id_string)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		err = controller.taskUseCase.DeleteMember(taskId, user_id, *user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	}

}
