package repository

import (
	"rastebazaar/pkg/domain/database"
	"rastebazaar/pkg/infra"

	"go.mongodb.org/mongo-driver/mongo"
)

// MarketRepo implements models.UserRepository
type MarketRepo struct {
	db *mongo.Client
}

// NewMarketRepo ..
func NewMarketRepo(infra *infra.Infrastructure) *MarketRepo {
	return &MarketRepo{
		db: infra.MongoClient,
	}
}

// MarketerLogin ...
func (r *MarketRepo) MarketerLogin(phone string) (string, error) {

	return "", nil
}

// MarketerVerify ...
func (r *MarketRepo) MarketerVerify(phone string, code string) (*database.Market, error) {
	return nil, nil
}
