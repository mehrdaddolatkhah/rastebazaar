package repository

import (
	"log"
	"rastebazaar/pkg/domain/database"
	"rastebazaar/pkg/infra"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

// CategoryRepo implements models.UserRepository
type CategoryRepo struct {
	db *mongo.Client
}

// NewCategoryRepo ..
func NewCategoryRepo(infra *infra.Infrastructure) *CategoryRepo {
	return &CategoryRepo{
		db: infra.MongoClient,
	}
}

// GetCategories ...
func (r *CategoryRepo) GetCategories() ([]database.Category, error) {

	// Pass these options to the Find method
	findOptions := options.Find()
	//findOptions.SetLimit(2)
	categories := []database.Category{}
	categoryCollection := r.db.Database(DatabaseName).Collection(database.CATEGORY_COLLECTION)

	cur, err := categoryCollection.Find(context.TODO(), bson.D{{}}, findOptions)

	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var category database.Category
		err := cur.Decode(&category)
		if err != nil {
			log.Fatal(err)
		}

		categories = append(categories, category)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return categories, nil
}

// AddCategory ...
func (r *CategoryRepo) AddCategory(category *database.Category) (*mongo.InsertOneResult, error) {

	categoryCollection := r.db.Database(DatabaseName).Collection(database.CATEGORY_COLLECTION)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, _ := categoryCollection.InsertOne(ctx, category)
	return result, nil
}
