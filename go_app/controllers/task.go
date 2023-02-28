package controllers

import (
	"encoding/json"
	"time"

	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/zimnushka/task_me_go/go_app/models"
	usecases "github.com/zimnushka/task_me_go/go_app/use_cases"
)

type TaskController struct {
	authUseCase usecases.AuthUseCase
	taskUseCase usecases.TaskUseCase
	corsUseCase usecases.CorsUseCase
	models.Controller
}

func (controller TaskController) Init(handler chi.Mux) models.Controller {
	controller.Url = "/task/"
	controller.RegisterController("", controller.taskHandler, handler)
	return controller.Controller
}

func (controller TaskController) taskHandler(w http.ResponseWriter, r *http.Request) {
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
			task, err := controller.taskUseCase.GetTaskById(id, *user.Id)
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
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonData))
	case "POST":
		var task models.Task
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		task.StartDate = time.Now().Format(time.RFC3339)
		newtask, err := controller.taskUseCase.AddTask(task, *user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		jsonData, err := json.Marshal(newtask)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonData))
	case "PUT":
		var task models.Task

		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = controller.taskUseCase.UpdateTask(task, *user.Id)
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
		err = controller.taskUseCase.DeleteTask(id, *user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	}

}
