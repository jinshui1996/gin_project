package main

import (
	"gin_project/router"

	"github.com/gin-gonic/gin"
)

func main() {
    var gin = gin.Default()
	router.LoadLoginRoutes(gin)
	router.LoadExampleRoutes(gin)
	gin.Run(":8080")
}