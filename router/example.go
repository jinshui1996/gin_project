package router

import (
	"github.com/gin-gonic/gin"
	"gin_project/controller"
	"gin_project/middlewares"
)

var exampleController = new(controller.ExampleController)

func LoadExampleRoutes(r *gin.Engine) *gin.RouterGroup {

	example := r.Group("/examples")

	example.Use(middlewares.Jwt()) // 需要鉴权的接口

	{
		example.GET("/test", func(c *gin.Context) {
		
					c.JSON(200, gin.H{
					"message": "example test",
				})
			
		})
		example.GET("/tsetCache", exampleController.TestCache)
		example.GET("/TestConsistentHash", exampleController.TestConsistentHash)
	}
	return example
}
