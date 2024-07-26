package main

import (
	"gin_project/router"
	"gin_project/config"

	"github.com/gin-gonic/gin"
)

func main() {
    var gin = gin.Default()
	router.LoadLoginRoutes(gin)
	router.LoadExampleRoutes(gin)
	var port = config.EnvConfig.Server.Port
	gin.Run(":" + port)
}