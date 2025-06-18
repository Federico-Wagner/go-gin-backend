package controllers

import (
	"backend_gin/config"
	"backend_gin/dto"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllProducts(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := config.ProductoCollection.Find(ctx, bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error buscando productos"})
		return
	}

	var productos []dto.Producto
	if err = cursor.All(ctx, &productos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error leyendo productos"})
		return
	}

	c.JSON(http.StatusOK, productos)
}