package controllers

import (
	"encoding/json"
	"time"

	"net/http"
	"strconv"
	"strings"

	"github.com/zimnushka/task_me_go/go_app/models"
	usecases "github.com/zimnushka/task_me_go/go_app/use_cases"
)

type TimeIntervalController struct {
	authUseCase     usecases.AuthUseCase
	intervalUseCase usecases.TimeIntervalUseCase
	corsUseCase     usecases.CorsUseCase
	models.Controller
}

func (controller TimeIntervalController) Init() models.Controller {
	controller.Url = "/timeIntervals/"
	controller.RegisterController("", controller.taskHandler)
	return controller.Controller
}

func (controller TimeIntervalController) taskHandler(w http.ResponseWriter, r *http.Request) {
	controller.corsUseCase.DisableCors(&w, r)
	user, err := controller.authUseCase.CheckToken(r.Header.Get(models.HeaderAuth))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		var jsonData []byte
		idString := strings.TrimPrefix(r.URL.Path, controller.Url)
		id, err := strconv.Atoi(idString)
		if err == nil {
			task, err := controller.intervalUseCase.GetIntervalsByTask(id, *user.Id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			jsonData, err = json.Marshal(task)

		} else {
			err = nil
			task, err := controller.intervalUseCase.GetIntervalsByUser(*user.Id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			jsonData, err = json.Marshal(task)

		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonData))
	case "POST":
		var jsonData []byte
		idString := strings.TrimPrefix(r.URL.Path, controller.Url)
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		var interval models.Interval
		interval.TaskId = id
		interval.UserId = *user.Id
		interval.TimeStart = time.Now().String()

		newInterval, err := controller.intervalUseCase.AddInterval(interval, interval.UserId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		jsonData, err = json.Marshal(newInterval)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonData))
	case "PUT":
		var jsonData []byte

		interval, err := controller.intervalUseCase.GetNotEndedInterval(*user.Id)
		interval.TimeEnd = time.Now().String()

		err = controller.intervalUseCase.UpdateInterval(*interval, *user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		jsonData, err = json.Marshal(interval)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonData))
	}

}
