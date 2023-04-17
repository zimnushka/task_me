package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zimnushka/task_me_go/go_app/models"
	usecases "github.com/zimnushka/task_me_go/go_app/use_cases"
)

type loginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

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
// @Param			new_user	body		models.User			true	"New User"
// @Success		200		{string}	string			"apiKey"
// @Router			/auth/registration [post]
func (controller AuthController) registrationHandler(c *gin.Context) {

	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return

	}

	newUser.Color = 4283658239
	apiKey, err := controller.authUseCase.Register(newUser)
	if err != nil {
		err.Call(c)
		return
	}

	c.String(http.StatusCreated, apiKey)

}

// @Summary		Login
// @Description	Login
// @ID				auth-login
// @Tags Auth
// @Accept			json
// @Produce		json
// @Param			login_user	body		loginParams			true	"Login user"
// @Success		200		{string}	string			"apiKey"
// @Router			/auth/login [post]
func (controller AuthController) loginHandler(c *gin.Context) {
	var params loginParams

	if err := c.BindJSON(&params); err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	apiKey, err := controller.authUseCase.Login(params.Email, params.Password)
	if err != nil {
		err.Call(c)
		return
	}
	c.String(http.StatusCreated, apiKey)

}
