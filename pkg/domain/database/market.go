package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const MARKET_COLLECTION string = "market"

// const (
//     REGISTERED Status = iota
//     VERIFIED
// )

type Market struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Phone      string             `bson:"phone,omitempty" json:"phone"`
	Username   string             `bson:"username,omitempty" json:"username"`
	FirstName  string             `bson:"first_name,omitempty" json:"first_name"`
	LastName   string             `bson:"last_name,omitempty" json:"last_name"`
	Status     string             `bson:"status,omitempty" json:"status"`
	MarketName string             `bson:"market_name,omitempty" json:"market_name"`
}
