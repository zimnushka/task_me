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

// @Summary		Get members
// @Description	Get project members
// @ID				project-get-members
// @Tags Project
// @Accept			json
// @Produce		json
// @Param			id	path		int		true	"Project id"
// @Success		200		{object}	[]models.User
// @Router			/project/member/{id} [get]
func (controller ProjectController) getProjectMembers(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	items, err := controller.projectUseCase.GetProjectUsers(id, *user.Id)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, items)
}

// @Summary		Add member
// @Description	Add project member
// @ID				project-add-members
// @Tags Project
// @Accept			json
// @Produce		json
// @Param			id	path		int		true	"Project id"
// @Param        email    query     string  false  "email new member"  Format(email)
// @Success		200		{string}	string
// @Router			/project/member/{id} [post]
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
	c.String(http.StatusOK, "")
}

// @Summary		Delete member
// @Description	Delete project members
// @ID				project-delete-members
// @Tags Project
// @Accept			json
// @Produce		json
// @Param			id	path		int		true	"Project id"
// @Param        userId    query     string  false  "User id for delete"
// @Success		200		{string}	string
// @Router			/project/member/{id} [put]
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

// @Summary		Get project by ID
// @Description	Get project by ID
// @ID				project-get-by-id
// @Tags Project
// @Accept			json
// @Produce		json
// @Param			id	path		int		true	"Project id"
// @Success		200		{object}	models.Project
// @Router			/project/{id} [get]
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

// @Summary		Get projects
// @Description	Get projects
// @ID				project-get
// @Tags Project
// @Accept			json
// @Produce		json
// @Success		200		{object}	[]models.Project
// @Router			/project [get]
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

// @Summary		Add project
// @Description	Add project
// @ID				project-add
// @Tags Project
// @Accept			json
// @Produce		json
// @Param			new_project	body		models.Project			true	"New project"
// @Success		200		{object}	models.Project
// @Router			/project [post]
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

// @Summary		Edit project
// @Description	Edit project
// @ID				project-edit
// @Tags Project
// @Accept			json
// @Produce		json
// @Param			project	body		models.Project			true	"Edit project"
// @Success		200		{object}	models.Project
// @Router			/project [put]
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

// @Summary		Delete project
// @Description	Delete project
// @ID				project-delete
// @Tags Project
// @Accept			json
// @Produce		json
// @Param			id	path		int		true	"Project id"
// @Success		200		{string}	string
// @Router			/project/{id} [delete]
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
