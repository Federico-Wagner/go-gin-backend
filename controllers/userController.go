package controllers

import (
	"backend_gin/config"
	"backend_gin/dto"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllUsers(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Obtenes un contexto con un timeout de 5 segundos
	defer cancel() 	// defer ejecuta la funcion cancel() al terminar la ejecucion de la funcion invocante para liberar recursos

	cursor, err := config.UsuarioCollection.Find(ctx, bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error buscando usuarios"})
		return
	}

	var usuarios []dto.Usuario
	if err = cursor.All(ctx, &usuarios); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error leyendo usuarios"})
		return
	}

	c.JSON(http.StatusOK, usuarios)
}

func CreateUser(c *gin.Context) {
	var user dto.Usuario
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := config.UsuarioCollection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error insertando usuario"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"inserted_id": result.InsertedID})
}

// GET /usuarios/:id
func GetUserByID(c *gin.Context) {
	idParam := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var usuario dto.Usuario
	err = config.UsuarioCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&usuario)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

// DELETE /usuarios/:id
func DeleteUserByID(c *gin.Context) {
	idParam := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := config.UsuarioCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado"})
}
