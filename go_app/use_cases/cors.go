package usecases

import (
	"net/http"
)

type CorsUseCase struct {
}

func (useCase *CorsUseCase) DisableCors(w *http.ResponseWriter, r *http.Request) {
	allowedHeaders := "Access-Control-Allow-Origin, Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-Requested-With, Authorization,X-CSRF-Token"

	(*w).Header().Add("Access-Control-Allow-Origin", "*")
	(*w).Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	(*w).Header().Add("Access-Control-Allow-Headers", allowedHeaders)
	if r.Method == "OPTIONS" {
		(*w).WriteHeader(http.StatusOK)
		return
	}
	if r.Method == "PATCH" {
		(*w).WriteHeader(http.StatusOK)
		return
	}

	return

}
