package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zimnushka/task_me_go/go_app/models"
	"github.com/zimnushka/task_me_go/go_app/repositories"
)

func UserControllerInit() models.Controller {
	controller := models.Controller{Url: "/user"}
	controller.RegisterController("", userController)
	return controller
}

func userController(w http.ResponseWriter, r *http.Request) {
	// GET, POST, PUT, DELETE
	switch r.Method {
	case "GET":
		users := repositories.GetUsers()
		s, _ := json.Marshal(users)
		fmt.Fprintf(w, string(s))
	case "POST":
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		repositories.AddUser(user)
		s, _ := json.Marshal(user)
		fmt.Fprintf(w, string(s))
	case "PUT":
		var user models.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		repositories.UpdateUser(user)

		// Do something with the Person struct...
		fmt.Fprintf(w, "User: %+v", user)
	case "DELETE":
		var user models.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		repositories.DeleteUser(*user.Id)
		fmt.Fprintf(w, "User: %+v", user)
	}

}
