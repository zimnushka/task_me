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

type TaskController struct {
	authUseCase usecases.AuthUseCase
	taskUseCase usecases.TaskUseCase
	models.Controller
}

func (controller TaskController) Init() models.Controller {
	controller.Url = "/task"
	controller.RegisterController("", controller.taskHandler)
	controller.RegisterController("/members/", controller.taskMemberHandler)
	return controller.Controller
}

func (controller TaskController) taskHandler(w http.ResponseWriter, r *http.Request) {
	user, err := controller.authUseCase.CheckToken(r.Header.Get(models.HeaderAuth))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			jsonData, err = json.Marshal(task)

		} else {
			err = nil
			tasks, err := controller.taskUseCase.GetAllTasks(*user.Id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			jsonData, err = json.Marshal(tasks)

		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, string(jsonData))
	case "POST":
		var task models.Task
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newtask, err := controller.taskUseCase.AddTask(task, *user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		s, err := json.Marshal(newtask)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, string(s))
	case "PUT":
		var task models.Task

		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = controller.taskUseCase.UpdateTask(task, *user.Id)
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
		err = controller.taskUseCase.DeleteTask(id, *user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "")
	}

}

func (controller TaskController) taskMemberHandler(w http.ResponseWriter, r *http.Request) {
	user, err := controller.authUseCase.CheckToken(r.Header.Get(models.HeaderAuth))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	idString := strings.Split(r.URL.Path, "/")
	taskId, err := strconv.Atoi(idString[len(idString)-1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		items, err := controller.taskUseCase.GetTaskUsers(taskId, *user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data, err := json.Marshal(items)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, string(data))
	case "PUT":
		var task models.TaskUser
		task.TaskId = taskId
		task.UserId, err = strconv.Atoi(r.URL.Query().Get("userId"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = controller.taskUseCase.AddMemberToTask(task.TaskId, task.UserId, *user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "")
	case "DELETE":
		var task models.TaskUser
		task.TaskId = taskId
		task.UserId, err = strconv.Atoi(r.URL.Query().Get("userId"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = controller.taskUseCase.DeleteMemberFromTask(task.TaskId, task.UserId, *user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "")
	}

}
