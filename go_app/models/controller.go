package models

import (
	"net/http"

	"github.com/go-chi/chi"
)

const HeaderAuth = "Authorization"

type Controller struct {
	Url string
}

func (c *Controller) RegisterController(subUrl string, handlerFunc func(http.ResponseWriter, *http.Request), handler chi.Mux) {
	handler.HandleFunc(c.Url+subUrl, handlerFunc)
}
