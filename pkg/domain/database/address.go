package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const ADDRESS_COLLECTION string = "address"

type Address struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID       primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	MarketId     primitive.ObjectID `bson:"market_id,omitempty" json:"market_id"`
	CountryName  string             `bson:"country_name,omitempty" json:"country_name"`
	ProvinceName string             `bson:"province_name,omitempty" json:"province_name"`
	CityName     string             `bson:"city_name,omitempty" json:"city_name"`
	Address      string             `bson:"address,omitempty" json:"address"`
	LandPhone    string             `bson:"land_phone,omitempty" json:"land_phone"`
	lat          string             `bson:"lat,omitempty" json:"lat"`
	lng          string             `bson:"lng,omitempty" json:"lng"`
}
