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

type TaskMemberController struct {
	authUseCase usecases.AuthUseCase
	taskUseCase usecases.TaskUseCase
	models.Controller
}

func (controller TaskMemberController) Init() models.Controller {
	controller.Url = "/taskMembers/"
	controller.RegisterController("", controller.taskMemberHandler)
	return controller.Controller
}

func (controller TaskMemberController) taskMemberHandler(w http.ResponseWriter, r *http.Request) {
	user, err := controller.authUseCase.CheckToken(r.Header.Get(models.HeaderAuth))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
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
		items, err := controller.taskUseCase.GetTaskUsers(taskId, *user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		data, err := json.Marshal(items)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, string(data))
	case "PUT":
		var task models.TaskUser
		task.TaskId = taskId
		task.UserId, err = strconv.Atoi(r.URL.Query().Get("userId"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = controller.taskUseCase.AddMemberToTask(task.TaskId, task.UserId, *user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "")
	case "DELETE":
		var task models.TaskUser
		task.TaskId = taskId
		task.UserId, err = strconv.Atoi(r.URL.Query().Get("userId"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = controller.taskUseCase.DeleteMemberFromTask(task.TaskId, task.UserId, *user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "")
	}

}
