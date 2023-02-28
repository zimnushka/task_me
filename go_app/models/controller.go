package models

import (
	"github.com/go-chi/chi"
)

const HeaderAuth = "Authorization"

type Controller interface {
	Init(handler chi.Mux)
}
