package main

import (
	"backend_gin/config"
	"backend_gin/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectToMongoDB()

	// Usuarios
	r.GET("/usuarios", controllers.GetAllUsers)
	r.POST("/usuarios", controllers.CreateUser)
	r.GET("/usuarios/:id", controllers.GetUserByID)
	r.DELETE("/usuarios/:id", controllers.DeleteUserByID)

	// Productos
	r.GET("/productos", controllers.GetAllProducts)

	r.Run(":8080")
}
