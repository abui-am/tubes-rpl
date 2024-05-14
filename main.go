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

	r.Use(corsMiddleware()) // Untuk mengatasi CORS
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/roles", controllers.CreateRole)
	r.GET("/roles", controllers.GetRoles)
	r.PUT("/roles/:id", controllers.UpdateRole)
	r.DELETE("/roles/:id", controllers.DeleteRoles)
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequiredAuth, controllers.Validate)
	r.GET("/items", middleware.RequiredAuth, controllers.GetItems)
	r.GET("/items/:id", middleware.RequiredAuth, controllers.GetItem) // GET: -> yang ngambil dari DB
	r.POST("/items", middleware.RequiredAuth, controllers.CreateItem)
	r.PUT("/items/:id", middleware.RequiredAuth, controllers.UpdateItem)      // PUT -> yang update ke DB
	r.DELETE("/items/:id", middleware.RequiredAuth, controllers.DeleteItem)   // PUT -> yang update ke DB
	r.POST("/borrowers", middleware.RequiredAuth, controllers.CreateBorrower) // POST -> yang masukin DB
	r.GET("/borrowers", middleware.RequiredAuth, controllers.GetBorrowers)    // GET: -> yang ngambil dari DB
	r.GET("/borrowers/:id", middleware.RequiredAuth, controllers.GetBorrower)
	r.PUT("/borrowers/:id", middleware.RequiredAuth, controllers.UpdateBorrower)
	r.DELETE("/borrowers/:id", middleware.RequiredAuth, controllers.DeleteBorrower)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
