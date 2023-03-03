package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zimnushka/task_me_go/go_app/models"
	usecases "github.com/zimnushka/task_me_go/go_app/use_cases"
)

type AuthController struct {
	authUseCase usecases.AuthUseCase
}

func (controller AuthController) Init(router *gin.Engine) {
	router.POST("/auth/login", controller.loginHandler)
	router.POST("/auth/registration", controller.registrationHandler)

}

// @Summary		Register
// @Description	Register new user
// @ID				auth-register
// @Tags Auth
// @Accept			json
// @Produce		json
// @Param			some_id	path		int				true	"Some ID"
// @Param			some_id	body		models.User			true	"Some ID"
// @Success		200		{string}	string			"apiKey"
// @Router			/auth/registration [get]
func (controller AuthController) registrationHandler(c *gin.Context) {

	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		return // TODO add error message
	}

	newUser.Color = 4283658239
	apiKey, err := controller.authUseCase.Register(newUser)
	if err != nil {
		return // TODO add error message
	}

	c.String(http.StatusCreated, apiKey)

}
func (controller AuthController) loginHandler(c *gin.Context) {

	type loginParams struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var params loginParams

	if err := c.BindJSON(&params); err != nil {
		return // TODO add error message
	}

	apiKey, err := controller.authUseCase.Login(params.Email, params.Password)
	if err != nil {
		return // TODO add error message
	}
	c.String(http.StatusCreated, apiKey)

}
