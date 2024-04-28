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
	initializers.DBSeed()
}

func main() {
	r := gin.Default()

	r.Use(corsMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/roles", controllers.GetRoles)
	r.GET("/users", controllers.GetUsers)
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequiredAuth, controllers.Validate)
	r.GET("/items", middleware.RequiredAuth, controllers.GetItems)
	r.GET("/items/:id", middleware.RequiredAuth, controllers.GetItem)
	r.POST("/items", middleware.RequiredAuth, controllers.CreateItem)
	r.PUT("/items/:id", middleware.RequiredAuth, controllers.UpdateItem)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
