package services

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

// New - initializes a new MongoDB client for use in the services package
func New(mongo *mongo.Client) {
	client = mongo
}

// returnCollectionPointer - returns a pointer to a MongoDB collection
func returnCollectionPointer(collection string) *mongo.Collection {
	return client.Database("pokedex-db").Collection(collection)
}
