package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type PokemonType struct {
	Slot int     `json:"slot"`
	Type Pokemon `json:"type"`
}

type PokemonTypes struct {
	Name        string        `json:"name"`
	Slot        int           `json:"slot"`
	PokemonType []PokemonType `json:"types"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonResponse struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Pokemon `json:"results"`
}

func fetchAllPokemon(limit string) ([]string, error) {
	url := "https://pokeapi.co/api/v2/pokemon?limit=" + limit
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error fetching data:", err)
		return nil, err
	}
	defer resp.Body.Close()

	var response PokemonResponse

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&response)

	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	var pokemonNames []string
	for _, pokemon := range response.Results {
		pokemonNames = append(pokemonNames, pokemon.Name)
	}
	return pokemonNames, nil
}

func fetchPokemonTypes(name string) (PokemonTypes, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name
	fmt.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error fetching data:", err)
		return PokemonTypes{}, err
	}
	defer resp.Body.Close()

	var response PokemonTypes
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&response)

	return PokemonTypes{Name: response.Name, PokemonType: response.PokemonType}, nil
}

// func fetchAllData(pokemons []string) []PokemonTypes {
// 	var data []PokemonTypes
// 	for i := 0; i < len(pokemons); i++ {
// 		pokemonName := pokemons[i]
// 		pokemon, _ := fetchPokemonTypes(pokemonName)
// 		data = append(data, pokemon)
// 	}
// 	return data
// }

func fetchAllData(pokemons []string) []PokemonTypes {
	var wg sync.WaitGroup
	var data []PokemonTypes
	for i := 0; i < len(pokemons); i++ {
		wg.Add(1)
		pokemonName := pokemons[i]
		go func(pokemonName string) {
			defer wg.Done()
			pokemon, _ := fetchPokemonTypes(pokemonName)
			data = append(data, pokemon)
		}(pokemonName)
	}
	wg.Wait()
	return data
}

func main() {
	start := time.Now()

	pokemons, err := fetchAllPokemon("100")
	if err != nil {
		return
	}
	fmt.Println(pokemons)
	data := fetchAllData(pokemons)
	fmt.Println(data)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
