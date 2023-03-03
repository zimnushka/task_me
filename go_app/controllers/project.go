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
	userUseCase    usecases.UserUseCase

	models.Controller
}

func (controller ProjectController) Init(router *gin.Engine) {
	router.GET("/project", controller.getProjects)
	router.GET("/project/:id", controller.getProjectById)
	router.POST("/project", controller.createProject)
	router.PUT("/project", controller.editProject)
	router.DELETE("/project/:id", controller.deleteProject)

	router.GET("/project/member/:id", controller.getProjectMembers)
	router.POST("/project/member/:id", controller.addProjectMember)
	router.PUT("/project/member/:id", controller.deleteProjectMember)
}

func (controller ProjectController) getProjectMembers(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		return // TODO add error message
	}
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return // TODO add error message
	}
	items, err := controller.projectUseCase.GetProjectUsers(id, *user.Id)
	if err != nil {
		return // TODO add error message
	}
	c.IndentedJSON(http.StatusOK, items)
}

func (controller ProjectController) addProjectMember(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		return // TODO add error message
	}
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return // TODO add error message
	}
	email := c.Query("email")

	member, err := controller.userUseCase.GetUserByEmail(email)
	if err != nil {
		return // TODO add error message
	}

	err = controller.projectUseCase.AddMemberToProject(id, *member.Id, *user.Id)
	if err != nil {
		return // TODO add error message
	}
}

func (controller ProjectController) deleteProjectMember(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		return // TODO add error message
	}
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return // TODO add error message
	}

	memberId, err := strconv.Atoi(c.Query("userId"))

	if err != nil {
		return // TODO add error message
	}

	if err = controller.projectUseCase.DeleteMemberFromProject(id, memberId, *user.Id); err != nil {
		return // TODO add error message
	}
	c.String(http.StatusOK, "")
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
