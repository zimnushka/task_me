package main

import (
	"net/http"

	"github.com/zimnushka/task_me_go/go_app/controllers"
)

func main() {
	controllers.AuthController{}.Init()
	controllers.UserController{}.Init()
	controllers.ProjectController{}.Init()
	controllers.TaskController{}.Init()
	controllers.ProjectMemberController{}.Init()
	controllers.TaskMemberController{}.Init()
	http.ListenAndServe(":8080", nil)
}
