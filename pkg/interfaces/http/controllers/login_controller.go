package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	domainrepository "github.com/williamokano/golang-web-api/pkg/app/domain/repository"
	"github.com/williamokano/golang-web-api/pkg/infrastructure/repository"
)

type LoginController struct {
	loginRepository domainrepository.LoginRepository
}

func NewLoginController(db *sql.DB) *LoginController {
	return &LoginController{
		loginRepository: repository.NewLoginRepository(db),
	}
}

func (controller *LoginController) GetLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"WHHAT": "hello",
	})
}

func (controller *LoginController) GetLoginByUsername(c *gin.Context) {
	username := c.Param("username")

	user, err := controller.loginRepository.FindByUsername(username)
	if err != nil {
		log.Println("failed")
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func (controller *LoginController) DoLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"login": "success",
	})
}
