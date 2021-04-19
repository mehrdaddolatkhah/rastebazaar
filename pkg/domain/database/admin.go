package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const ADMIN_COLLECTION string = "admin"

type Admin struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Mobile    string             `bson:"phone,omitempty" json:"phone"`
	Firstname string             `bson:"firstname,omitempty" json:"first_name"`
	Lastname  string             `bson:"lastname,omitempty" json:"last_name"`
	Password  string             `bson:"password,omitempty" json:"password"`
}
