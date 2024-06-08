package routes

import (
	"golang-gin3/config"
	"golang-gin3/middleware"
	"golang-gin3/src/controller"
	"golang-gin3/src/repository"
	"golang-gin3/src/service"

	"github.com/gin-gonic/gin"
)

func OrderRouter(r *gin.RouterGroup) {
	api := r.Group("/orders")

	orderRepository := repository.NewRepositoryOrder(config.Db)
	orderService := service.NewOrderService(orderRepository)
	orderController := controller.NewOrderController(orderService)

	api.Use(middleware.Authorization())
	api.POST("/", orderController.Create)
	api.GET("/", orderController.FindAll)
	api.GET("/:id", orderController.FindById)
	api.PUT("/:id", orderController.Update)
	api.DELETE("/:id", orderController.Delete)
}
