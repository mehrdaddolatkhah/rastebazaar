package database

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// const (
//     REGISTERED Status = iota
//     VERIFIED
// )

const USER_COLLECTION string = "user"

// User defines the storage form of a beer
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Phone     string             `bson:"phone,omitempty" json:"phone"`
	Firstname string             `bson:"firstname,omitempty" json:"first_name"`
	Lastname  string             `bson:"lastname,omitempty" json:"last_name"`
	Email     string             `bson:"email,omitempty" json:"email"`
	BirthDay  time.Time          `bson:"birthday,omitempty" json:"birthday"`
	Identify  string             `bson:"identify,omitempty" json:"identify"`
	Password  string             `bson:"password,omitempty" json:"password"`
	Suspend   bool               `bson:"suspend,omitempty" json:"suspend"`
	Status    string             `bson:"status,omitempty" json:"status"`
}
