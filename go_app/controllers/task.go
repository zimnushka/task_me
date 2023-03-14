package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zimnushka/task_me_go/go_app/models"
	usecases "github.com/zimnushka/task_me_go/go_app/use_cases"
)

type TaskController struct {
	authUseCase usecases.AuthUseCase
	taskUseCase usecases.TaskUseCase
}

func (controller TaskController) Init(router *gin.Engine) {
	router.GET("/task/project/:id", controller.getTaskByProject)

	router.GET("/task", controller.getUserTasks)
	router.GET("/task/:id", controller.getTaskById)
	router.POST("/task", controller.addTask)
	router.PUT("/task", controller.editTask)
	router.DELETE("/task/:id", controller.deleteTask)

	router.GET("/task/member/:id", controller.getTaskMembers)
	router.POST("/task/member/:id", controller.editTaskMembersList)
}

func (controller TaskController) getTaskById(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		return // TODO add error message
	}
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return // TODO add error message
	}
	item, err := controller.taskUseCase.GetTaskById(id, *user.Id)
	if err != nil {
		return // TODO add error message
	}
	c.IndentedJSON(http.StatusOK, item)
}

func (controller TaskController) getUserTasks(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		return // TODO add error message
	}
	items, err := controller.taskUseCase.GetAllTasks(*user.Id)
	if err != nil {
		return // TODO add error message
	}
	c.IndentedJSON(http.StatusOK, items)
}

func (controller TaskController) addTask(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		return // TODO add error message
	}
	var item models.Task
	if err := c.BindJSON(&item); err != nil {
		return // TODO add error message
	}
	newItem, err := controller.taskUseCase.AddTask(item, *user.Id)
	if err != nil {
		return // TODO add error message
	}

	c.IndentedJSON(http.StatusOK, newItem)
}

func (controller TaskController) editTask(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		return // TODO add error message
	}
	var item models.Task
	if err := c.BindJSON(&item); err != nil {
		return // TODO add error message
	}
	if err := controller.taskUseCase.UpdateTask(item, *user.Id); err != nil {
		return // TODO add error message
	}

	c.IndentedJSON(http.StatusOK, item)
}

func (controller TaskController) deleteTask(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		return // TODO add error message
	}
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return // TODO add error message
	}
	if err = controller.taskUseCase.DeleteTask(id, *user.Id); err != nil {
		return // TODO add error message
	}

	c.String(http.StatusOK, "")
}

func (controller TaskController) getTaskMembers(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		return // TODO add error message
	}
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return // TODO add error message
	}
	items, err := controller.taskUseCase.GetMembers(id, *user.Id)
	if err != nil {
		return // TODO add error message
	}
	c.IndentedJSON(http.StatusOK, items)
}

func (controller TaskController) editTaskMembersList(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		return // TODO add error message
	}
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return // TODO add error message
	}
	var items []models.User
	if err := c.BindJSON(&items); err != nil {
		return // TODO add error message
	}
	err = controller.taskUseCase.UpdateMembersList(id, items, *user.Id)
	if err != nil {
		return // TODO add error message
	}
	c.String(http.StatusOK, "")
}

func (controller TaskController) getTaskByProject(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		return // TODO add error message
	}
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return // TODO add error message
	}
	item, err := controller.taskUseCase.GetTaskByProjectId(id, *user.Id)
	if err != nil {
		return // TODO add error message
	}
	c.IndentedJSON(http.StatusOK, item)
}
