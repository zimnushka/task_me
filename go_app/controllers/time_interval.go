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

func (controller TimeIntervalController) getIntervalsByTask(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		return // TODO add error message
	}
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return // TODO add error message
	}
	items, err := controller.intervalUseCase.GetIntervalsByTask(id, *user.Id)
	if err != nil {
		return // TODO add error message
	}
	c.IndentedJSON(http.StatusOK, items)
}

func (controller TimeIntervalController) getIntervalsByUser(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		return // TODO add error message
	}

	items, err := controller.intervalUseCase.GetIntervalsByUser(*user.Id)
	if err != nil {
		return // TODO add error message
	}
	c.IndentedJSON(http.StatusOK, items)
}

func (controller TimeIntervalController) AddInterval(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		return // TODO add error message
	}
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return // TODO add error message
	}
	item, err := controller.intervalUseCase.AddInterval(id, *user.Id)
	if err != nil {
		return // TODO add error message
	}
	c.IndentedJSON(http.StatusOK, item)
}

func (controller TimeIntervalController) FinishInterval(c *gin.Context) {
	user, err := controller.authUseCase.CheckToken(c.GetHeader(models.HeaderAuth))
	if err != nil {
		return // TODO add error message
	}
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return // TODO add error message
	}

	if err := controller.intervalUseCase.FinishInterval(id, *user.Id); err != nil {
		return // TODO add error message
	}
	c.String(http.StatusOK, "")
}
