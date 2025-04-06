package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type Pokemon struct {
	NationalID   int         `json:"nationalID,omitempty" bson:"_id,omitempty"`
	Name         string      `json:"name,omitempty" bson:"name,omitempty"`
	Type         []string    `json:"type,omitempty" bson:"type,omitempty"`
	Abilities    []Abilities `json:"abilities,omitempty" bson:"abilities,omitempty"`
	PokeDexEntry string      `json:"pokedexEntry,omitempty" bson:"pokedexEntry,omitempty"`
	Form         string      `json:"form,omitempty" bson:"form,omitempty"`
}

type Abilities struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Hidden      bool   `json:"hidden,omitempty" bson:"hidden,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
}

func (p *Pokemon) GetAllPokemons() ([]Pokemon, error) {
	collection := returnCollectionPointer("pokemon")

	var pokemons []Pokemon

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var pokemon Pokemon
		cursor.Decode(&pokemon)
		pokemons = append(pokemons, pokemon)
	}

	return pokemons, nil
}

func (p *Pokemon) InsertPokemon(pokemon Pokemon) error {
	collection := returnCollectionPointer("pokemon")
	_, err := collection.InsertOne(context.TODO(), Pokemon{
		NationalID:   pokemon.NationalID,
		Name:         pokemon.Name,
		Type:         pokemon.Type,
		Abilities:    pokemon.Abilities,
		PokeDexEntry: pokemon.PokeDexEntry,
		Form:         pokemon.Form,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (p *Pokemon) GetPokemonByName(name string) (Pokemon, error) {
	collection := returnCollectionPointer("pokemon")

	var pokemon Pokemon

	err := collection.FindOne(context.Background(), bson.D{{"name", name}}).Decode(&pokemon)
	if err != nil {
		log.Println(err)
		return Pokemon{}, err
	}
	return pokemon, nil
}

func (p *Pokemon) DeletePokemonByName(name string) error {
	collection := returnCollectionPointer("pokemon")
	_, err := collection.DeleteOne(context.Background(), bson.D{{"name", name}})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
