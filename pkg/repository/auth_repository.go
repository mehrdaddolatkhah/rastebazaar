package repository

import (
	"rastebazaar/pkg/domain/database"

	"rastebazaar/pkg/infra"

	"go.mongodb.org/mongo-driver/mongo"
)

// AuthRepo implements models.UserRepository
type AuthRepo struct {
	db *mongo.Client
}

// NewAuthRepo ..
func NewAuthRepo(infra *infra.Infrastructure) *AuthRepo {
	return &AuthRepo{
		db: infra.MongoClient,
	}
}

// Login ..
func (r *AuthRepo) Login(phone string) (string, error) {
	return "", nil
}

// Verify ..
func (r *AuthRepo) Verify(phone string, code string) (*database.User, error) {
	return nil, nil
}
