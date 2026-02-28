package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//user model struct

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email"  json:"email"`
	Password string             `bson:"password"  json:"password"`
	IsAdmin  bool               `bson:"is_admin"   json:"is_admin"`
}

type Product struct {
	ID          primitive.ObjectID `bson:"id_omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Price       float64            `bson:"price" json:"price"`
	Description string             `bson:"description" json:"description"`
	Image       string             `bson:"image" json:"image"`
}
