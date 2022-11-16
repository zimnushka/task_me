package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zimnushka/task_me_go/go_app/models"
	usecases "github.com/zimnushka/task_me_go/go_app/use_cases"
)

type AuthController struct {
	authUseCase usecases.AuthUseCase
	models.Controller
}

func (controller AuthController) Init() models.Controller {
	controller.Url = "/auth"
	controller.RegisterController("/login", controller.loginHandler)
	controller.RegisterController("/registration", controller.registrationHandler)
	return controller.Controller
}

func (controller AuthController) registrationHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response, err := controller.authUseCase.Register(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, response)
	}

}
func (controller AuthController) loginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		type loginParams struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		var params loginParams
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response, err := controller.authUseCase.Login(params.Email, params.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, response)
	}

}
