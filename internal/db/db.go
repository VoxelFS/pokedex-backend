package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

// ConnectToMongo - connects to the MongoDB database
func ConnectToMongo() (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://pokedex-db.nsvzzdu.mongodb.net/?retryWrites=true&w=majority&appName=pokedex-db").SetServerAPIOptions(serverAPI)

	username := os.Getenv("MONGO_DB_USERNAME")
	password := os.Getenv("MONGO_DB_PASSWORD")

	opts.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}

	log.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return client, nil
}
