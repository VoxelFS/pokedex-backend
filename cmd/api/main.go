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

func main() {
	mongoClient, err := db.ConnectToMongo()
	if err != nil {
		log.Panic(err)
	}

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
