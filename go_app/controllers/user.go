package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zimnushka/task_me_go/go_app/app"
	"github.com/zimnushka/task_me_go/go_app/models"
	usecases "github.com/zimnushka/task_me_go/go_app/use_cases"
)

type UserController struct {
	userUseCase usecases.UserUseCase
	authUseCase usecases.AuthUseCase
}

func (controller UserController) Init(router *gin.Engine) {
	router.GET("/user", controller.getUsers)
	router.GET("/user/:id", controller.getUserById)
	router.GET("/user/me", controller.getUserMe)
	router.POST("/user", controller.addUser)
	router.PUT("/user", controller.editUser)
	router.DELETE("/user/:id", controller.deleteUser)
}

func (controller UserController) getUserMe(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		err.Call(c)
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (controller UserController) getUserById(c *gin.Context) {
	_, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth)) // TODO add check permission for this
	if err != nil {
		err.Call(c)
		return
	}
	idString := c.Param("id")
	id, strToIntErr := strconv.Atoi(idString)
	if strToIntErr != nil {
		app.AppErrorByError(strToIntErr).Call(c)
		return
	}
	item, err := controller.userUseCase.GetUserById(id)
	if err != nil {
		err.Call(c)
		return
	}
	c.IndentedJSON(http.StatusOK, item)
}

func (controller UserController) getUsers(c *gin.Context) {
	_, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth)) // TODO add check permission for this
	if err != nil {
		err.Call(c)
		return
	}
	items, err := controller.userUseCase.GetAllUsers()
	if err != nil {
		err.Call(c)
		return
	}
	c.IndentedJSON(http.StatusOK, items)
}

func (controller UserController) addUser(c *gin.Context) {
	_, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		err.Call(c)
		return
	}
	var item models.User
	if err := c.BindJSON(&item); err != nil {
		app.AppErrorByError(err).Call(c)
		return
	}
	newItem, err := controller.userUseCase.AddUser(item)
	if err != nil {
		err.Call(c)
		return
	}

	c.IndentedJSON(http.StatusOK, newItem)
}

func (controller UserController) editUser(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		err.Call(c)
		return
	}
	var item models.User
	if err := c.BindJSON(&item); err != nil {
		app.AppErrorByError(err).Call(c)
		return
	}
	if _, err := controller.userUseCase.UpdateUser(item, *user.Id); err != nil {
		err.Call(c)
		return
	}

	c.IndentedJSON(http.StatusOK, item)
}

func (controller UserController) deleteUser(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		err.Call(c)
		return
	}
	idString := c.Param("id")
	id, strToIntErr := strconv.Atoi(idString)
	if strToIntErr != nil {
		app.AppErrorByError(strToIntErr).Call(c)
		return
	}
	if err = controller.userUseCase.DeleteUser(id, *user.Id); err != nil {
		err.Call(c)
		return
	}

	c.String(http.StatusOK, "")
}
