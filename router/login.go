package router

import (
	"github.com/gin-gonic/gin"
	"gin_project/controller"
)

func LoadLoginRoutes(r *gin.Engine) *gin.RouterGroup {
	login := r.Group("/login")

	{
		login.GET("/oneTest", func(c *gin.Context) {
	
			name := c.Query("name")
			password := c.Query("password")
			tokenStr, err := controller.Login(name, password)
			if err != nil {
				c.JSON(500, gin.H{
					"message": err.Error(),
				})
			} else {
				c.JSON(200, gin.H{
					"token": tokenStr,
				})
			}
		})
	}
	return login
}