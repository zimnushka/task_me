package main

import (
	"net/http"

	"github.com/zimnushka/task_me_go/go_app/controllers"
)

func main() {
	authController := controllers.AuthControllerInit()
	authController.RegisterController()

	http.ListenAndServe(":8080", nil)
}
