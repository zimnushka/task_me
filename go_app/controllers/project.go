package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zimnushka/task_me_go/go_app/models"
	usecases "github.com/zimnushka/task_me_go/go_app/use_cases"
)

type ProjectController struct {
	authUseCase    usecases.AuthUseCase
	projectUseCase usecases.ProjectUseCase

	models.Controller
}

func (controller ProjectController) Init(router *gin.Engine) {
	router.GET("/project", controller.getProjects)
	router.GET("/project/:id", controller.getProjectById)
	router.POST("/project", controller.createProject)
	router.PUT("/project", controller.editProject)
	router.DELETE("/project/:id", controller.deleteProject)
}

func (controller ProjectController) getProjectById(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		return // TODO add error message
	}
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return // TODO add error message
	}
	item, err := controller.projectUseCase.GetProjectById(id, *user.Id)
	if err != nil {
		return // TODO add error message
	}
	c.IndentedJSON(http.StatusOK, item)
}

func (controller ProjectController) getProjects(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		return // TODO add error message
	}
	items, err := controller.projectUseCase.GetAllProjects(*user.Id)
	if err != nil {
		return // TODO add error message
	}
	c.IndentedJSON(http.StatusOK, items)
}

func (controller ProjectController) createProject(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		return // TODO add error message
	}

	var project models.Project
	if err := c.BindJSON(&project); err != nil {
		return // TODO add error message
	}
	newproject, err := controller.projectUseCase.AddProject(project, *user.Id)
	if err != nil {
		return // TODO add error message
	}
	c.IndentedJSON(http.StatusOK, *newproject)
}

func (controller ProjectController) editProject(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		return // TODO add error message
	}
	var project models.Project
	if err := c.BindJSON(&project); err != nil {
		return // TODO add error message
	}
	if err := controller.projectUseCase.UpdateProject(project, *user.Id); err != nil {
		return // TODO add error message
	}

	c.IndentedJSON(http.StatusOK, project)
}

func (controller ProjectController) deleteProject(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		return // TODO add error message
	}
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return // TODO add error message
	}
	if err = controller.projectUseCase.DeleteProject(id, *user.Id); err != nil {
		return // TODO add error message
	}

	c.String(http.StatusOK, "")
}
