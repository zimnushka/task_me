package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zimnushka/task_me_go/go_app/models"
	usecases "github.com/zimnushka/task_me_go/go_app/use_cases"
)

func AuthControllerInit() models.Controller {
	controller := models.Controller{Url: "/auth"}
	controller.RegisterController("/login", loginController)
	controller.RegisterController("/registration", registrationController)
	return controller
}

func registrationController(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var user models.User
		var useCase usecases.AuthUseCase
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response, err := useCase.Register(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, response)
	}

}
func loginController(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		type loginParams struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		var params loginParams
		var useCase usecases.AuthUseCase
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response, err := useCase.Login(params.Email, params.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, response)
	}

}
