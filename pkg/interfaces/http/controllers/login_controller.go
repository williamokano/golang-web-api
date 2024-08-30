package controllers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	db *sql.DB
}

func NewLoginController(db *sql.DB) *LoginController {
	return &LoginController{
		db: db,
	}
}

func (controller *LoginController) GetLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"WHHAT": "hello",
	})
}

func (controller *LoginController) DoLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"login": "success",
	})
}
