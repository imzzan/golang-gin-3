package main

import (
	"fmt"
	"golang-gin3/config"
	routes "golang-gin3/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load func env
	config.LoadConfig()
	config.LoadDb()

	router := gin.Default()
	api := router.Group("/api")

	routes.UserRouter(api)
	routes.ProductRouter(api)
	routes.TweetRoutes(api)
	routes.OrderRouter(api)

	api.GET("/debug/pprof", gin.WrapH(router))
	port := fmt.Sprintf(":%v", config.ENV.PORT)
	router.Run(port)
}
