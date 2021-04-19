package domain

import (
	"rastebazaar/pkg/domain/database"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
}

type ProductRepository interface {
}

type MarketRepository interface {
	MarketerLogin(phone string) (string, error)
	MarketerVerify(phone string, code string) (*database.Market, error)
}

type AdminRepository interface {
	GetAdminByID(userID string) (database.Admin, error)
	AdminLogin(mobile string, password string) (database.Admin, error)
	AdminRegister(*database.Admin) (*mongo.InsertOneResult, error)
}

type CategoryRepository interface {
	GetCategories() ([]database.Category, error)
	AddCategory(*database.Category) (*mongo.InsertOneResult, error)
}

type AuthRepository interface {
	Login(phone string) (string, error)
	Verify(phone string, code string) (*database.User, error)
}
