package web

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/williamokano/golang-web-api/docs"
	taskusecase "github.com/williamokano/golang-web-api/internal/application/task"
	"github.com/williamokano/golang-web-api/internal/infrastructure/infrastructure/persistence"
	"github.com/williamokano/golang-web-api/pkg/http/task"
)

type WebApplication struct {
	db     *sql.DB
	router *gin.Engine
}

func NewServer(db *sql.DB) *WebApplication {
	return &WebApplication{
		db:     db,
		router: gin.Default(),
	}
}

func (app WebApplication) SetupServer() *http.Server {

	app.setupTaskHandler()

	// Setup swagger
	docs.SwaggerInfo.BasePath = "/"
	app.router.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	app.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server := &http.Server{
		Addr:    ":8080",
		Handler: app.router.Handler(),
	}

	return server
}

func (app WebApplication) setupTaskHandler() {
	taskRepository := persistence.NewTaskRepository(app.db)

	// use cases
	getAllTasksUC := taskusecase.NewGetAllTasks(taskRepository)
	createTaskUC := taskusecase.NewCreateTask(taskRepository)

	// handler
	taskHandler := task.NewTaskHandler(getAllTasksUC, createTaskUC)

	// setup routes
	task.SetupRoutes(app.router, taskHandler)
}
