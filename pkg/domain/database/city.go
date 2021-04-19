package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CITY_COLLECTION string = "city"

type City struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name,omitempty" json:"name"`
	Province string             `bson:"province,omitempty" json:"province"`
	Country  string             `bson:"country,omitempty" json:"country"`
}
