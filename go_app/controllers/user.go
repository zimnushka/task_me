package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
		return // TODO add error message
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (controller UserController) getUserById(c *gin.Context) {
	_, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth)) // TODO add check permission for this
	if err != nil {
		return // TODO add error message
	}
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return // TODO add error message
	}
	item, err := controller.userUseCase.GetUserById(id)
	if err != nil {
		return // TODO add error message
	}
	c.IndentedJSON(http.StatusOK, item)
}

func (controller UserController) getUsers(c *gin.Context) {
	_, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth)) // TODO add check permission for this
	if err != nil {
		return // TODO add error message
	}
	items, err := controller.userUseCase.GetAllUsers()
	if err != nil {
		return // TODO add error message
	}
	c.IndentedJSON(http.StatusOK, items)
}

func (controller UserController) addUser(c *gin.Context) {
	_, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth)) // TODO add check permission for this
	if err != nil {
		return // TODO add error message
	}
	var item models.User
	if err := c.BindJSON(&item); err != nil {
		return // TODO add error message
	}
	newItem, err := controller.userUseCase.AddUser(item)
	if err != nil {
		return // TODO add error message
	}

	c.IndentedJSON(http.StatusOK, newItem)
}

func (controller UserController) editUser(c *gin.Context) {
	_, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth)) // TODO add check permission for this
	if err != nil {
		return // TODO add error message
	}
	var item models.User
	if err := c.BindJSON(&item); err != nil {
		return // TODO add error message
	}
	if _, err := controller.userUseCase.UpdateUser(item); err != nil {
		return // TODO add error message
	}

	c.IndentedJSON(http.StatusOK, item)
}

func (controller UserController) deleteUser(c *gin.Context) {
	_, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth)) // TODO add check permission for this
	if err != nil {
		return // TODO add error message
	}
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return // TODO add error message
	}
	if err = controller.userUseCase.DeleteUser(id); err != nil {
		return // TODO add error message
	}

	c.String(http.StatusOK, "")
}
