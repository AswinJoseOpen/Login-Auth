package server

import (
	"github.com/AswinJoseOpen/Login-Auth/config"
	"github.com/AswinJoseOpen/Login-Auth/controller"
	"github.com/AswinJoseOpen/Login-Auth/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Start(db *gorm.DB, config config.AppConfig) {
	router := gin.Default()
	testController := controller.NewController(db, config)
	setupRoutes(router, testController, db, config)
	router.Run()

}
func setupRoutes(r *gin.Engine, testController *controller.Controller, db *gorm.DB, config config.AppConfig) {
	r.GET("/", middleware.AuthMiddleWare(db, config), testController.Test)
	r.POST("/signup", testController.SignUp)
	r.POST("/login", testController.Login)
}
