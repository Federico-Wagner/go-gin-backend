package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Usuario struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Nombre string             `json:"nombre" bson:"nombre"`
	Email  string             `json:"email" bson:"email"`
}