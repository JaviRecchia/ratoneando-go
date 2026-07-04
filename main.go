package main

import (
	"github.com/gin-gonic/gin"

	"ratoneando/config"
	"ratoneando/middlewares"
	"ratoneando/routes"
	"ratoneando/utils/cache"
	"ratoneando/utils/logger"
)

func main() {
	logger.Init()
	config.Init()
	cache.Init()

	if config.ENV == "production" || config.ENV == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else if config.ENV == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	port := config.PORT

	router := gin.Default()

	middlewares.CORS(router)

	// Register routes
	routes.RegisterRoutes(router)

	// Start the server
	logger.Log("Starting server on port " + port)
	if err := router.Run(":" + port); err != nil {
		logger.LogFatal("Could not start server: " + err.Error())
	}
}
