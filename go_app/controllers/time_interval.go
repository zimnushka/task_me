package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zimnushka/task_me_go/go_app/app"
	"github.com/zimnushka/task_me_go/go_app/models"
	usecases "github.com/zimnushka/task_me_go/go_app/use_cases"
)

type descriptionModel struct {
	Description string `json:"description"`
}

type TimeIntervalController struct {
	authUseCase     usecases.AuthUseCase
	intervalUseCase usecases.TimeIntervalUseCase
}

func (controller TimeIntervalController) Init(router *gin.Engine) {
	router.GET("/timeIntervals", controller.getIntervalsByUser)
	router.GET("/timeIntervals/task/:id", controller.getIntervalsByTask)
	router.GET("/timeIntervals/project/:id", controller.getIntervalsByProject)
	router.POST("/timeIntervals/:id", controller.AddInterval)
	router.PUT("/timeIntervals", controller.FinishInterval)
}

// @Summary		Get intervals by task ID
// @Description	Get intervals by task ID
// @ID				intervals-get-by-task-id
// @Tags Intervals
// @Accept			json
// @Produce		json
// @Param			id	path		int		true	"Task id"
// @Success		200		{object}	[]models.Interval
// @Router			/timeIntervals/task/{id} [get]
func (controller TimeIntervalController) getIntervalsByTask(c *gin.Context) {
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
	items, err := controller.intervalUseCase.GetIntervalsByTask(id, *user.Id)
	if err != nil {
		err.Call(c)
		return
	}
	c.IndentedJSON(http.StatusOK, items)
}

// @Summary		Get intervals by project ID
// @Description	Get intervals by project ID
// @ID				intervals-get-by-project-id
// @Tags Intervals
// @Accept			json
// @Produce		json
// @Param			id	path		int		true	"Project id"
// @Success		200		{object}	[]models.Interval
// @Router			/timeIntervals/project/{id} [get]
func (controller TimeIntervalController) getIntervalsByProject(c *gin.Context) {
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
	items, err := controller.intervalUseCase.GetIntervalsByProject(id, *user.Id)
	if err != nil {
		err.Call(c)
		return
	}
	c.IndentedJSON(http.StatusOK, items)
}

// @Summary		Get my intervals
// @Description	Get my intervals
// @ID				intervals-get-my
// @Tags Intervals
// @Accept			json
// @Produce		json
// @Success		200		{object}	[]models.Interval
// @Router			/timeIntervals [get]
func (controller TimeIntervalController) getIntervalsByUser(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		err.Call(c)
		return
	}

	items, err := controller.intervalUseCase.GetIntervalsByUser(*user.Id)
	if err != nil {
		err.Call(c)
		return
	}
	c.IndentedJSON(http.StatusOK, items)
}

// @Summary		Start interval
// @Description	Start interval
// @ID				intervals-start
// @Tags Intervals
// @Accept			json
// @Produce		json
// @Param			id	path		int		true	"Task id"
// @Success		200		{object}	models.Interval
// @Router			/timeIntervals/{id} [post]
func (controller TimeIntervalController) AddInterval(c *gin.Context) {
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
	item, err := controller.intervalUseCase.AddInterval(id, *user)
	if err != nil {
		err.Call(c)
		return
	}
	c.IndentedJSON(http.StatusOK, item)
}

// @Summary		Stop interval
// @Description	Stop interval
// @ID				intervals-stop
// @Tags Intervals
// @Accept			json
// @Produce		json
// @Param			id	path		int		true	"Task id"
// @Success		200		{object}	models.Interval
// @Router			/timeIntervals/{id} [put]
func (controller TimeIntervalController) FinishInterval(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		err.Call(c)
		return
	}

	var item descriptionModel
	if err := c.BindJSON(&item); err != nil {
		item.Description = ""
	}

	if err := controller.intervalUseCase.FinishInterval(*user.Id, item.Description); err != nil {
		err.Call(c)
		return
	}
	c.String(http.StatusOK, "")
}
