package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zimnushka/task_me_go/go_app/models"
	usecases "github.com/zimnushka/task_me_go/go_app/use_cases"
)

type TimeIntervalController struct {
	authUseCase     usecases.AuthUseCase
	intervalUseCase usecases.TimeIntervalUseCase
}

func (controller TimeIntervalController) Init(router *gin.Engine) {
	router.GET("/timeIntervals", controller.getIntervalsByUser)
	router.GET("/timeIntervals/:id", controller.getIntervalsByTask)
	router.POST("/timeIntervals/:id", controller.AddInterval)
	router.PUT("/timeIntervals/:id", controller.FinishInterval)
}

// @Summary		Get intervals by task ID
// @Description	Get intervals by task ID
// @ID				intervals-get-by-task-id
// @Tags Intervals
// @Accept			json
// @Produce		json
// @Param			id	path		int		true	"Task id"
// @Success		200		{object}	[]models.Interval
// @Router			/timeIntervals/{id} [get]
func (controller TimeIntervalController) getIntervalsByTask(c *gin.Context) {
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
	items, err := controller.intervalUseCase.GetIntervalsByTask(id, *user.Id)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
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
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	items, err := controller.intervalUseCase.GetIntervalsByUser(*user.Id)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
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
	item, err := controller.intervalUseCase.AddInterval(id, *user.Id)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
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

	if err := controller.intervalUseCase.FinishInterval(id, *user.Id); err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, "")
}
