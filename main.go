package main

import (
	"backend_gin/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {

		var pingMsg map[string]string = controllers.GetPing()
		fmt.Print(pingMsg)
		c.JSON(200, pingMsg)
	})

	r.Run() // Por defecto corre en :8080
}
