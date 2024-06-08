package routes

import (
	"golang-gin3/config"
	"golang-gin3/middleware"
	"golang-gin3/src/controller"
	"golang-gin3/src/repository"
	"golang-gin3/src/service"

	"github.com/gin-gonic/gin"
)

func ProductRouter(r *gin.RouterGroup) {
	api := r.Group("/products")

	productRepository := repository.NewProductRepository(config.Db)
	userRepository := repository.NewUserRepository(config.Db)
	productService := service.NewProductService(productRepository, userRepository)
	productController := controller.NewProductController(productService)

	api.GET("/", productController.GetAllProducts)
	api.Use(middleware.Authorization())
	api.POST("/", productController.Create)
	api.GET("/me", productController.FindByUserId)
	api.GET("/:id", middleware.JustMe(), productController.FindById)
	api.PUT("/:id", middleware.JustMe(), productController.Update)
	api.DELETE("/:id", middleware.JustMe(), productController.Delete)
}
