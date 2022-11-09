package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zimnushka/task_me_go/go_app/models"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

func AuthControllerInit() models.Controller {
	controller := models.Controller{Url: "/auth"}
	controller.RegisterController("/login", loginController)
	controller.RegisterController("/registration", registrationController)
	return controller
}

func registrationController(w http.ResponseWriter, r *http.Request) {
	// GET, POST, PUT, DELETE

	if r.Method == "POST" {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		repositories.AddUser(user)
		s, _ := json.Marshal(user)
		fmt.Fprintf(w, string(s))
	}

}
func loginController(w http.ResponseWriter, r *http.Request) {
	// GET, POST, PUT, DELETE

	if r.Method == "POST" {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		repositories.AddUser(user)
		s, _ := json.Marshal(user)
		fmt.Fprintf(w, string(s))
	}

}
