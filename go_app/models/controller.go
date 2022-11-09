package models

import (
	"net/http"
)

type Controller struct {
	Url        string
	HandleFunc func(http.ResponseWriter, *http.Request)
}

func (c *Controller) RegisterController() {
	http.HandleFunc(c.Url, c.HandleFunc)
}
