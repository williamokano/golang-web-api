package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/williamokano/golang-web-api/pkg/interfaces/http/controllers"
	"github.com/williamokano/golang-web-api/pkg/interfaces/middleware"
)

func (app *Application) setupRootRoutes() {
	app.router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	loginController := controllers.NewLoginController(app.db)
	app.router.GET("/login", loginController.GetLogin)
	app.router.POST("/login", loginController.DoLogin)
}

func (app *Application) SetupServer() *http.Server {
	app.router.Use(middleware.MetricMiddleware())

	app.setupRootRoutes()

	server := &http.Server{
		Addr:    ":8080",
		Handler: app.router.Handler(),
	}

	return server
}
