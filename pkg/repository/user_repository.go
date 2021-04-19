package repository

import (
	"rastebazaar/pkg/infra"

	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepo implements models.UserRepository
type UserRepo struct {
	db *mongo.Client
}

// NewUserRepo ..
func NewUserRepo(infra *infra.Infrastructure) *UserRepo {
	return &UserRepo{
		db: infra.MongoClient,
	}
}
