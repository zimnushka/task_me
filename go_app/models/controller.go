package models

import (
	"net/http"
)

const HeaderAuth = "Authorization"

type Controller struct {
	Url string
}

func (c *Controller) RegisterController(subUrl string, handlerFunc func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(c.Url+subUrl, handlerFunc)
}
