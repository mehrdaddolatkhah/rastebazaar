package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const PRODUCT_COLLECTION string = "product"

type Product struct {
	ID               primitive.ObjectID     `bson:"_id,omitempty" json:"id"`
	Name             string                 `bson:"name,omitempty" json:"name"`
	EnName           string                 `bson:"en_name,omitempty" json:"en_name"`
	Description      string                 `bson:"description,omitempty" json:"description"`
	ShortDescription string                 `bson:"short_description,omitempty" json:"short_description"`
	Category         []string               `bson:"category,omitempty" json:"category"`
	Attributes       map[string]interface{} `bson:"attributes,omitempty" json:"attributes"`
	Images           []string               `bson:"images,omitempty" json:"image"`
	Thumbnails       []string               `bson:"thumbnails,omitempty" json:"thumbnail"`
	Videos           []string               `bson:"videos,omitempty" json:"video"`
	Suspend          bool                   `bson:"suspend,omitempty" json:"suspend"`
}
