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

type ProjectMemberController struct {
	authUseCase    usecases.AuthUseCase
	projectUseCase usecases.ProjectUseCase
	userUseCase    usecases.UserUseCase
	corsUseCase    usecases.CorsUseCase
	models.Controller
}

func (controller ProjectMemberController) Init() models.Controller {
	controller.Url = "/projectMembers/"
	controller.RegisterController("", controller.projectMemberHandler)
	return controller.Controller
}

func (controller ProjectMemberController) projectMemberHandler(w http.ResponseWriter, r *http.Request) {
	controller.corsUseCase.DisableCors(w, r)
	user, err := controller.authUseCase.CheckToken(r.Header.Get(models.HeaderAuth))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	idString := strings.Split(r.URL.Path, "/")
	projectId, err := strconv.Atoi(idString[len(idString)-1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		items, err := controller.projectUseCase.GetProjectUsers(projectId, *user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		data, err := json.Marshal(items)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(data))
	case "PUT":
		var project models.ProjectUser
		project.ProjectId = projectId
		email := r.URL.Query().Get("email")
		member, err := controller.userUseCase.GetUserByEmail(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		err = controller.projectUseCase.AddMemberToProject(project.ProjectId, *member.Id, *user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "")
	case "DELETE":
		var project models.ProjectUser
		project.ProjectId = projectId
		project.UserId, err = strconv.Atoi(r.URL.Query().Get("userId"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = controller.projectUseCase.DeleteMemberFromProject(project.ProjectId, project.UserId, *user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "")
	}

}
