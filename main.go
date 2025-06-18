package main

import (
	"backend_gin/controllers"
	"backend_gin/dto"

	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		var pingMsg map[string]string = controllers.GetPing()
		fmt.Print(pingMsg)
		c.JSON(200, pingMsg)
	})

	// GET /usuarios - devuelve todos los usuarios
	router.GET("/usuarios", func(c *gin.Context) {
		var usuarios []dto.Usuario = controllers.GetAllUsers()
		c.JSON(http.StatusOK, usuarios)
	})

	// POST /usuarios - crea un nuevo usuario
	router.POST("/usuario", func(c *gin.Context) {
		var userDTO dto.Usuario
		// Bind del JSON recibido al struct Usuario
		if err := c.BindJSON(&userDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var newUser = controllers.CreateNewUser(userDTO)

		c.JSON(http.StatusCreated, newUser)
	})

	router.DELETE("/usuario", func(c *gin.Context) {
		var userDTO dto.Usuario
		// Bind del JSON recibido al struct Usuario
		if err := c.BindJSON(&userDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var success = controllers.DeleteUser(userDTO)

		c.JSON(http.StatusCreated, success)
	})

	router.PUT("/usuario", func(c *gin.Context) {
		var userDTO dto.Usuario
		// Bind del JSON recibido al struct Usuario
		if err := c.BindJSON(&userDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var success = controllers.EditUser(userDTO)

		c.JSON(http.StatusCreated, success)
	})

	router.Run() // Por defecto corre en :8080
}
