package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CATEGORY_COLLECTION string = "category"

type Category struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"ـid"`
	Name   string             `bson:"name,omitempty" json:"name"`
	Parent string             `bson:"parent,omitempty" json:"parent"`
}
