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
	return models.Controller{Url: "/auth", HandleFunc: authController}
}

func authController(w http.ResponseWriter, r *http.Request) {
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
