package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UsuarioCollection *mongo.Collection
var ProductoCollection *mongo.Collection

func ConnectToMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// URI de conexión local
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("❌ Error conectando a Mongo:", err)
	}

	// Testear conexión
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("❌ Mongo no responde:", err)
	}

	UsuarioCollection = client.Database("usersMS").Collection("usuarios")
	ProductoCollection = client.Database("usersMS").Collection("productos")
	log.Println("✅ Conectado a Mongo en localhost")
}