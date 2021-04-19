package infra

import (
	"html/template"

	"go.mongodb.org/mongo-driver/mongo"
)

type Infrastructure struct {
	MongoClient *mongo.Client
	Template    *template.Template
}

type BackgroundVars struct {
	UserID string
}
