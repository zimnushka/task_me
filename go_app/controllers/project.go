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
	controller.Url = "/project"
	controller.IsNeedAuth = true
	controller.RegisterController("", controller.projectHandler)
	controller.RegisterController("/members/", controller.projectMemberHandler)
	return controller.Controller
}

func (controller ProjectController) projectHandler(w http.ResponseWriter, r *http.Request) {
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
			project, err := controller.projectUseCase.GetProjectById(id, *user.Id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			jsonData, err = json.Marshal(project)

		} else {
			err = nil
			projects, err := controller.projectUseCase.GetAllProjects(*user.Id)
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
		newproject, err := controller.projectUseCase.AddProject(project, *user.Id)
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
		err = controller.projectUseCase.UpdateProject(project, *user.Id)
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
		err = controller.projectUseCase.DeleteProject(id, *user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "")
	}

}

func (controller ProjectController) projectMemberHandler(w http.ResponseWriter, r *http.Request) {
	user, err := controller.authUseCase.CheckToken(r.Header.Get(models.HeaderAuth))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	idString := strings.Split(r.URL.Path, "/")
	projectId, err := strconv.Atoi(idString[len(idString)-1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		items, err := controller.projectUseCase.GetProjectUsers(projectId, *user.Id)
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
		var project models.ProjectUser
		project.ProjectId = projectId
		project.UserId, err = strconv.Atoi(r.URL.Query().Get("userId"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = controller.projectUseCase.AddMemberToProject(project.ProjectId, project.UserId, *user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "")
	case "DELETE":
		var project models.ProjectUser
		project.ProjectId = projectId
		project.UserId, err = strconv.Atoi(r.URL.Query().Get("userId"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = controller.projectUseCase.DeleteMemberFromProject(project.ProjectId, project.UserId, *user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "")
	}

}
