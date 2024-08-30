package web

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Application struct {
	db     *sql.DB
	router *gin.Engine
}

func NewServer(db *sql.DB) *Application {
	return &Application{
		db:     db,
		router: gin.Default(),
	}
}
