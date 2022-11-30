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
	authUseCase    usecases.AuthUseCase
	projectUseCase usecases.ProjectUseCase
	models.Controller
}

func (controller ProjectController) Init() models.Controller {
	controller.Url = "/project/"
	controller.RegisterController("", controller.projectHandler)
	return controller.Controller
}

func (controller ProjectController) projectHandler(w http.ResponseWriter, r *http.Request) {
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
			project, err := controller.projectUseCase.GetProjectById(id, *user.Id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			jsonData, err = json.Marshal(project)

		} else {
			err = nil
			projects, err := controller.projectUseCase.GetAllProjects(*user.Id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			jsonData, err = json.Marshal(projects)

		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, string(jsonData))
	case "POST":
		var project models.Project
		err := json.NewDecoder(r.Body).Decode(&project)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		newproject, err := controller.projectUseCase.AddProject(project, *user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		s, err := json.Marshal(newproject)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, string(s))
	case "PUT":
		var project models.Project

		err := json.NewDecoder(r.Body).Decode(&project)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = controller.projectUseCase.UpdateProject(project, *user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		fmt.Fprintf(w, "")
	case "DELETE":
		idString := strings.TrimPrefix(r.URL.Path, controller.Url)
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = controller.projectUseCase.DeleteProject(id, *user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		fmt.Fprintf(w, "")
	}

}
