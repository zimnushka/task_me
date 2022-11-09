package main

import (
	"net/http"

	"github.com/zimnushka/task_me_go/go_app/controllers"
)

func main() {
	controllers.AuthControllerInit()
	controllers.UserControllerInit()

	http.ListenAndServe(":8080", nil)
}
