package main

import (
	"net/http"

	"github.com/zimnushka/task_me_go/go_app/controllers"
)

func main() {
	controllers.AuthController{}.Init()
	controllers.UserController{}.Init()
	http.ListenAndServe(":8080", nil)
}
