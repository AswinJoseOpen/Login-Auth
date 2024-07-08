package controller

import (
	"fmt"
	"net/http"

	"github.com/AswinJoseOpen/Login-Auth/config"
	"github.com/AswinJoseOpen/Login-Auth/model"
	"github.com/AswinJoseOpen/Login-Auth/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	service *service.ServiceImpl
}

func NewController(db *gorm.DB, config config.AppConfig) *Controller {
	testController := &Controller{
		service: &service.ServiceImpl{
			DB:     db,
			Config: config,
		},
	}
	return testController
}

func (t *Controller) Test(c *gin.Context) {
	fmt.Println("Testing Controller")
	msg := t.service.TestService(c)
	c.JSON(http.StatusOK, msg)
}
func (t *Controller) SignUp(c *gin.Context) {
	var body model.Users
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	if err := t.service.SignUp(c, &body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	// respond
	c.JSON(http.StatusOK, gin.H{})
}

func (t *Controller) Login(c *gin.Context) {
	var body model.Users
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	tokenResponse, err := t.service.Login(c, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, tokenResponse)
}
