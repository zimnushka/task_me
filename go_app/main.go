package main

import (
	"net/http"

	"github.com/zimnushka/task_me_go/go_app/controllers"
)

func main() {
	http.HandleFunc("/users", controllers.UserController)
	http.ListenAndServe(":8080", nil)
}
