package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/williamokano/golang-web-api/internal/application/task"
)

type TaskHandler struct {
	getAllTasksUC *usecase.GetAllTasks
	createTaskUC  *usecase.CreateTask
}

func NewTaskHandler(getAllTasksUC *usecase.GetAllTasks, createTaskUC *usecase.CreateTask) *TaskHandler {
	return &TaskHandler{
		getAllTasksUC: getAllTasksUC,
		createTaskUC:  createTaskUC,
	}
}

// GetAll returns all the registered tasks
// @Summary Get all tasks
// @Description Get a list of all tasks that wasn't deleted
// @Tags tasks
// @Accept json
// @Success 200
// @Failure 501
// @Router /tasks [get]
func (h *TaskHandler) GetAll(c *gin.Context) {
	allTasks, err := h.getAllTasksUC.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, allTasks)
}

func (h *TaskHandler) Create(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "Not implemented"})
}
