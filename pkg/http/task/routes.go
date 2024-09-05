package task

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine, handler *TaskHandler) {
	taskGroup := router.Group("/tasks")

	taskGroup.GET("/", handler.GetAll)
	taskGroup.POST("/", handler.Create)
}
