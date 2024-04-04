package main

import (
	"gihub.com/abui-am/tubes-rpl/controllers"
	"gihub.com/abui-am/tubes-rpl/middleware"

	"gihub.com/abui-am/tubes-rpl/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequiredAuth, controllers.Validate)
	r.Run() // listen and serve on 0.0.0.0:8080
}
