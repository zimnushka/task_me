package usecases

import (
	"net/http"
)

type CorsUseCase struct {
}

func (useCase *CorsUseCase) DisableCors(w http.ResponseWriter, r *http.Request) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, X-Requested-With, Authorization,X-CSRF-Token"

	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	}

	return
}
