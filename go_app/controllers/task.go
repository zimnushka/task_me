package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zimnushka/task_me_go/go_app/app"
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

// @Summary		Get task by ID
// @Description	Get task by ID
// @ID				task-get-by-id
// @Tags Task
// @Accept			json
// @Produce		json
// @Param			id	path		int		true	"Task id"
// @Success		200		{object}	models.Task
// @Router			/task/{id} [get]
func (controller TaskController) getTaskById(c *gin.Context) {
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
	item, err := controller.taskUseCase.GetTaskById(id, *user.Id)
	if err != nil {
		err.Call(c)
		return
	}
	c.IndentedJSON(http.StatusOK, item)
}

// @Summary		Get tasks
// @Description	Get tasks
// @ID				task-get
// @Tags Task
// @Accept			json
// @Produce		json
// @Success		200		{object}	[]models.Task
// @Router			/task [get]
func (controller TaskController) getUserTasks(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		err.Call(c)
		return
	}
	items, err := controller.taskUseCase.GetAllTasks(*user.Id)
	if err != nil {
		err.Call(c)
		return
	}
	c.IndentedJSON(http.StatusOK, items)
}

// @Summary		Add task
// @Description	Add task
// @ID				task-add
// @Tags Task
// @Accept			json
// @Produce		json
// @Param			new_task	body		models.Task			true	"New task"
// @Success		200		{object}	models.Task
// @Router			/task [post]
func (controller TaskController) addTask(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		err.Call(c)
		return
	}
	var item models.Task
	if err := c.BindJSON(&item); err != nil {
		app.AppErrorByError(err).Call(c)
		return
	}
	newItem, err := controller.taskUseCase.AddTask(item, *user.Id)
	if err != nil {
		err.Call(c)
		return
	}

	c.IndentedJSON(http.StatusOK, newItem)
}

// @Summary		Edit task
// @Description	Edit task
// @ID				task-edit
// @Tags Task
// @Accept			json
// @Produce		json
// @Param			edit_task	body		models.Task			true	"Edit task"
// @Success		200		{object}	models.Task
// @Router			/task [put]
func (controller TaskController) editTask(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		err.Call(c)
		return
	}
	var item models.Task
	if err := c.BindJSON(&item); err != nil {
		app.AppErrorByError(err).Call(c)
		return
	}
	if err := controller.taskUseCase.UpdateTask(item, *user.Id); err != nil {
		err.Call(c)
		return
	}

	c.IndentedJSON(http.StatusOK, item)
}

// @Summary		Delete task
// @Description	Delete task
// @ID				task-delete
// @Tags Task
// @Accept			json
// @Produce		json
// @Param			id	path		int		true	"Task id"
// @Success		200		{string}	string
// @Router			/task/{id} [delete]
func (controller TaskController) deleteTask(c *gin.Context) {
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
	if err = controller.taskUseCase.DeleteTask(id, *user.Id); err != nil {
		err.Call(c)
		return
	}

	c.String(http.StatusOK, "")
}

// @Summary		Get task members
// @Description	Get task members
// @ID				task-get-members
// @Tags Task
// @Accept			json
// @Produce		json
// @Param			id	path		int		true	"Task id"
// @Success		200		{object}	models.User
// @Router			/task/member/{id} [get]
func (controller TaskController) getTaskMembers(c *gin.Context) {
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
	items, err := controller.taskUseCase.GetMembers(id, *user.Id)
	if err != nil {
		err.Call(c)
		return
	}
	c.IndentedJSON(http.StatusOK, items)
}

// @Summary		Add task members
// @Description	Add task members
// @ID				task-add-members
// @Tags Task
// @Accept			json
// @Produce		json
// @Param			update_members_list	body		[]models.User			true	"Update members list"
// @Success		200		{string}	string
// @Router			/task/member/{id} [post]
func (controller TaskController) editTaskMembersList(c *gin.Context) {
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
	var items []models.User
	if err := c.BindJSON(&items); err != nil {
		app.AppErrorByError(err).Call(c)
		return
	}
	err = controller.taskUseCase.UpdateMembersList(id, items, *user.Id)
	if err != nil {
		err.Call(c)
		return
	}
	c.String(http.StatusOK, "")
}

// @Summary		Get task by project
// @Description	Get task by project
// @ID				task-get-by-project
// @Tags Task
// @Accept			json
// @Produce		json
// @Param			id	path		int		true	"Project id"
// @Success		200		{object}	[]models.Task
// @Router			/task/project/{id} [get]
func (controller TaskController) getTaskByProject(c *gin.Context) {
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
	item, err := controller.taskUseCase.GetTaskByProjectId(id, *user.Id)
	if err != nil {
		err.Call(c)
		return
	}
	c.IndentedJSON(http.StatusOK, item)
}
