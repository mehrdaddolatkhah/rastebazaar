package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const COUNTRY_COLLECTION string = "country"

type Country struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name,omitempty" json:"name"`
	IsForeign bool               `bson:"is_foreign,omitempty" json:"is_foreign"`
}
