package models

import (
	"net/http"
)

type Controller struct {
	Url        string
	IsNeedAuth bool
}

func (c *Controller) RegisterController(subUrl string, handlerFunc func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(c.Url+subUrl, handlerFunc)
}
