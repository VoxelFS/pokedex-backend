package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/pokedex-backend/internal/services"
	"github.com/pokedex-backend/pkg/write_response"
	"log"
	"net/http"
)

var pokemon services.Pokemon

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

func GetPokemonByName(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "name")
	pokemon1, err := pokemon.GetPokemonByName(id)
	if err != nil {
		write_response.RequestErrorHandler(w, "No Pokemon Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pokemon1)
}

func DeletePokemonByName(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "name")
	err := pokemon.DeletePokemonByName(id)
	if err != nil {
		write_response.RequestErrorHandler(w, "No Pokemon Found", http.StatusNotFound)
		return
	}

	write_response.StatusOkHandler(w, "Pokemon Deleted Successfully")
}
