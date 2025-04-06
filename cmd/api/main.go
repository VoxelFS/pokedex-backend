package main

import (
	"context"
	"github.com/pokedex-backend/internal/db"
	"github.com/pokedex-backend/internal/handlers"
	"github.com/pokedex-backend/internal/services"
	"log"
	"net/http"
	"time"
)

// main - The entry point of the application
func main() {
	mongoClient, err := db.ConnectToMongo()
	if err != nil {
		log.Panic(err)
	}

	// Creates a 15-second context so if the connection with MongoDB takes longer than 15-seconds, it cancels the operation.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	services.New(mongoClient)

	// Sets up a server at port 8080
	log.Println("Server running in port", 8080)
	log.Fatal(http.ListenAndServe(":8080", handlers.CreateRouter()))

}
