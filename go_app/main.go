package main

import (
	"net/http"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/zimnushka/task_me_go/go_app/docs"

	"github.com/zimnushka/task_me_go/go_app/controllers"
	usecases "github.com/zimnushka/task_me_go/go_app/use_cases"
)

// @title           TaskMe API
// @version         1.0
// @description     Swagger documentation taskMe API

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	r := chi.NewRouter()

	controllers.AuthController{}.Init(*r)
	controllers.UserController{}.Init(*r)
	controllers.ProjectController{}.Init(*r)
	controllers.TaskController{}.Init(*r)
	controllers.TaskProjectController{}.Init(*r)
	controllers.ProjectMemberController{}.Init(*r)
	controllers.TaskMemberController{}.Init(*r)
	controllers.TimeIntervalController{}.Init(*r)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
	r.HandleFunc("/", handler)
	http.ListenAndServe(":8080", r)

}

func handler(w http.ResponseWriter, r *http.Request) {
	corsUseCase := usecases.CorsUseCase{}
	corsUseCase.DisableCors(&w, r)

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	}

}
