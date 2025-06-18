package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Producto struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Nombre string             `json:"nombre" bson:"nombre"`
	Categoria  string             `json:"categoria" bson:"categoria"`
	Precio  float64             `json:"precio" bson:"precio"`
}