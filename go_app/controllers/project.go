package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zimnushka/task_me_go/go_app/models"
	usecases "github.com/zimnushka/task_me_go/go_app/use_cases"
)

type ProjectController struct {
	authUseCase    usecases.AuthUseCase
	projectUseCase usecases.ProjectUseCase

	models.Controller
}

func (controller ProjectController) Init(router *gin.Engine) models.Controller {
	// controller.Url = "/project/"
	// controller.RegisterController("", controller.projectHandler, handler)
	return controller.Controller
}

func (controller ProjectController) projectHandler(w http.ResponseWriter, r *http.Request) {
	// controller.corsUseCase.DisableCors(&w, r) // TODO fix CORS
	user, err := controller.authUseCase.CheckToken(r.Header.Get(models.HeaderAuth))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		var jsonData []byte
		idString := "-1" //TODO
		id, err := strconv.Atoi(idString)
		if err == nil {
			project, err := controller.projectUseCase.GetProjectById(id, *user.Id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			jsonData, _ = json.Marshal(project)

		} else {
			err = nil
			projects, err := controller.projectUseCase.GetAllProjects(*user.Id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			jsonData, _ = json.Marshal(projects)

		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
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
		jsonData, err := json.Marshal(newproject)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
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
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	case "DELETE":
		idString := "-1" //TODO
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
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	}

}
