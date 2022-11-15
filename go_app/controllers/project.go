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

type ProjectController struct {
	projectUseCase usecases.ProjectUseCase
	models.Controller
}

func (controller ProjectController) Init() models.Controller {
	controller.Url = "/project"
	controller.IsNeedAuth = true
	controller.RegisterController("", controller.projectHandler)
	return controller.Controller
}

func (controller ProjectController) projectHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var jsonData []byte
		idString := strings.TrimPrefix(r.URL.Path, controller.Url)
		id, err := strconv.Atoi(idString)
		if err == nil {
			project, err := controller.projectUseCase.GetProjectById(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			jsonData, err = json.Marshal(project)

		} else {
			err = nil
			projects, err := controller.projectUseCase.GetAllProjects()
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			jsonData, err = json.Marshal(projects)

		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, string(jsonData))
	case "POST":
		var project models.Project
		err := json.NewDecoder(r.Body).Decode(&project)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newproject, err := controller.projectUseCase.AddProject(project)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		s, err := json.Marshal(newproject)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, string(s))
	case "PUT":
		var project models.Project

		err := json.NewDecoder(r.Body).Decode(&project)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = controller.projectUseCase.UpdateProject(project)
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
		err = controller.projectUseCase.DeleteProject(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "")
	}

}
