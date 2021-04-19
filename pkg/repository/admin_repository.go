package repository

import (
	"time"

	"rastebazaar/pkg/domain/database"

	"rastebazaar/pkg/infra"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

// AdminRepo implements models.UserRepository
type AdminRepo struct {
	db *mongo.Client
}

// NewAdminRepo ..
func NewAdminRepo(infra *infra.Infrastructure) *AdminRepo {
	return &AdminRepo{
		db: infra.MongoClient,
	}
}

// GetAdminByID ...
func (r *AdminRepo) GetAdminByID(userID string) (database.Admin, error) {

	objID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		panic(err)
	}

	admin := database.Admin{}
	adminCollection := r.db.Database(DatabaseName).Collection(database.ADMIN_COLLECTION)

	err = adminCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&admin)

	if err != nil {
		return admin, err
	}

	return admin, nil
}

// AdminRegister ...
func (r *AdminRepo) AdminRegister(admin *database.Admin) (*mongo.InsertOneResult, error) {

	adminCollection := r.db.Database(DatabaseName).Collection(database.ADMIN_COLLECTION)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, _ := adminCollection.InsertOne(ctx, admin)
	return result, nil
}

// AdminLogin ...
func (r *AdminRepo) AdminLogin(mobile string, password string) (database.Admin, error) {

	admin := database.Admin{}

	adminCollection := r.db.Database(DatabaseName).Collection(database.ADMIN_COLLECTION)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err := adminCollection.FindOne(ctx, bson.M{"phone": mobile, "password": password}).Decode(&admin)

	if err != nil {
		return admin, err
	}

	return admin, nil
}
