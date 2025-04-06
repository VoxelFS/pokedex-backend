package handlers

import (
	"github.com/go-chi/chi/v5"
	chimiddle "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/pokedex-backend/internal/middleware"
)

type Response struct {
	Message string
	Code    int
}

// CreateRouter - sets up and returns a configured chi router with routes and middleware
func CreateRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{""}, // set up allowed origins here
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}), chimiddle.StripSlashes)

	router.Route("/api", func(r chi.Router) {

		r.Get("/pokemons", GetAllPokemons)
		r.Get("/pokemon/{name}", GetPokemonByName)

		// Routes that require authorisation
		r.With(middleware.Authorisation).Group(func(r chi.Router) {
			r.Post("/addpokemon", InsertPokemon)
			r.Delete("/pokemon/{name}", DeletePokemonByName)
			r.Post("/logout", LogoutHandler)
		})

		r.Post("/login", LoginHandler)

	})
	return router
}
