package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const PROVINCE_COLLECTION string = "province"

type Province struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name,omitempty"`
	Country string             `bson:"country,omitempty"`
}
