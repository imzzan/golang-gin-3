package routes

import (
	"golang-gin3/config"
	"golang-gin3/middleware"
	"golang-gin3/src/controller"
	"golang-gin3/src/repository"
	"golang-gin3/src/service"

	"github.com/gin-gonic/gin"
)

func TweetRoutes(r *gin.RouterGroup) {
	api := r.Group("/tweets")

	tweetRepo := repository.NewTweetRepo(config.Db)
	tweetService := service.NewTweetService(tweetRepo)
	tweetController := controller.NewTweetController(tweetService)

	api.Use(middleware.Authorization())
	api.POST("/", tweetController.Create)
	api.GET("/", tweetController.FindAll)
	api.GET("/:id", tweetController.FindById)
	api.PUT("/:id", tweetController.Update)
	api.DELETE("/:id", tweetController.Delete)
}
