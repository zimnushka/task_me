package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/zimnushka/task_me_go/go_app/app"
	"github.com/zimnushka/task_me_go/go_app/controllers"
	_ "github.com/zimnushka/task_me_go/go_app/docs"
)

// @title           TaskMe API
// @version         1.0
// @description     Swagger documentation taskMe API

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {string} "test"
// @Router / [get]

func main() {
	app.SetConfig(app.ReleaseConfig)

	router := gin.Default()
	router.Use(CORSMiddleware())

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	//TODO (OPTIMIZATION): taskDTO to task controller
	//TODO (OPTIMIZATION): get all page info from one request
	//TODO (FIX): get all user permission fix

	controllers.AuthController{}.Init(router)
	controllers.UserController{}.Init(router)
	controllers.ProjectController{}.Init(router)
	controllers.TaskController{}.Init(router)
	controllers.TimeIntervalController{}.Init(router)

	router.Run(":8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
