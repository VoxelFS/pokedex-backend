package services

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func New(mongo *mongo.Client) {
	client = mongo
}

func returnCollectionPointer(collection string) *mongo.Collection {
	return client.Database("pokedex-db").Collection(collection)
}
