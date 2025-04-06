package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/pokedex-backend/internal/services"
	"github.com/pokedex-backend/pkg/write_response"
	"log"
	"net/http"
)

// pokemon is a variable that gives access to methods defined in the services.Pokemon service
var pokemon services.Pokemon

// GetAllPokemons handles GET requests to retrieve all Pokémon from the database
func GetAllPokemons(w http.ResponseWriter, r *http.Request) {
	pokemons, err := pokemon.GetAllPokemons()
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pokemons)
}

// InsertPokemon handles POST requests to add a new Pokémon to the database
func InsertPokemon(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&pokemon)
	if err != nil {
		log.Println(err)
		return
	}
	errP := pokemon.InsertPokemon(pokemon)
	if errP != nil {
		write_response.RequestErrorHandler(w, errP.Error(), http.StatusBadRequest)
		return
	}
	res := Response{
		Message: "Pokemon Added Successfully",
		Code:    http.StatusCreated,
	}

	jsonResponse, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(jsonResponse)
}

// GetPokemonByName handles GET requests to fetch a Pokémon by its name
func GetPokemonByName(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "name") // Extract Pokémon name from the URL
	pokemon1, err := pokemon.GetPokemonByName(id)
	if err != nil {
		write_response.RequestErrorHandler(w, "No Pokemon Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pokemon1)
}

// DeletePokemonByName handles DELETE requests to remove a Pokémon by name
func DeletePokemonByName(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "name") // Extract Pokémon name from the URL
	err := pokemon.DeletePokemonByName(id)
	if err != nil {
		write_response.RequestErrorHandler(w, "No Pokemon Found", http.StatusNotFound)
		return
	}

	write_response.StatusOkHandler(w, "Pokemon Deleted Successfully")
}
