package routes

import (
	"golang-gin3/config"
	"golang-gin3/middleware"
	"golang-gin3/src/controller"
	"golang-gin3/src/repository"
	"golang-gin3/src/service"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	api := r.Group("/users")

	userRepository := repository.NewUserRepository(config.Db)
	userServices := service.NewUserService(userRepository)
	userController := controller.UserController(userServices)

	api.POST("/register", userController.Register)
	api.POST("/login", userController.Login)
	api.GET("/me", middleware.Authorization(), userController.GetMe)
}
