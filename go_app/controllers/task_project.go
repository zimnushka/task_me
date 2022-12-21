package controllers

import (
	"encoding/json"

	"net/http"
	"strconv"
	"strings"

	"github.com/zimnushka/task_me_go/go_app/models"
	usecases "github.com/zimnushka/task_me_go/go_app/use_cases"
)

type TaskProjectController struct {
	authUseCase usecases.AuthUseCase
	taskUseCase usecases.TaskUseCase
	corsUseCase usecases.CorsUseCase
	models.Controller
}

func (controller TaskProjectController) Init() models.Controller {
	controller.Url = "/taskProject/"
	controller.RegisterController("", controller.taskHandler)
	return controller.Controller
}

func (controller TaskProjectController) taskHandler(w http.ResponseWriter, r *http.Request) {
	controller.corsUseCase.DisableCors(&w, r)

	user, err := controller.authUseCase.CheckToken(r.Header.Get(models.HeaderAuth))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		var jsonData []byte
		idString := strings.TrimPrefix(r.URL.Path, controller.Url)
		id, err := strconv.Atoi(idString)
		if err == nil {
			task, err := controller.taskUseCase.GetTaskByProjectId(id, *user.Id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			jsonData, err = json.Marshal(task)

		} else {
			err = nil
			tasks, err := controller.taskUseCase.GetAllTasks(*user.Id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			jsonData, err = json.Marshal(tasks)

		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonData))
	}

}
