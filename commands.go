package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/EveBisk/pokedex/internal/pokeapi"
)

func commandExit(reg *config, arg string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)

	return nil
}

func commandHelp(reg *config, arg string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	registry := getCommands()

	for _, val := range registry {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}

	return nil
}

func commandMap(cfg *config, arg string) error {
	url := pokeapi.GetLocationURL(cfg.nextLocationsURL)

	var locations_response pokeapi.LocationAreasResponse
	var err error

	entries, exist := cfg.pokeapiCache.Get(url)

	if exist {
		err = json.Unmarshal(entries, &locations_response)

	} else {
		locations_response, err = cfg.pokeapiClient.ListLocations(url)
	}

	if err != nil {
		return err
	}

	for _, entry := range locations_response.Results {
		fmt.Println(entry.Name)
	}

	cfg.prevLocationsURL = locations_response.Previous
	cfg.nextLocationsURL = locations_response.Next

	b, err := json.Marshal(locations_response)
	if err != nil {
		return err
	}

	cfg.pokeapiCache.Add(url, b)

	return nil
}

func commandMapBack(cfg *config, arg string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	var locations_response pokeapi.LocationAreasResponse
	var err error

	entries, exist := cfg.pokeapiCache.Get(*cfg.prevLocationsURL)

	if exist {
		err = json.Unmarshal(entries, &locations_response)
	} else {
		locations_response, err = cfg.pokeapiClient.ListLocations(*cfg.prevLocationsURL)
	}

	if err != nil {
		return err
	}

	for _, entry := range locations_response.Results {
		fmt.Println(entry.Name)
	}

	cfg.prevLocationsURL = locations_response.Previous
	cfg.nextLocationsURL = locations_response.Next

	return nil
}

func commandExplore(cfg *config, name string) error {
	fmt.Printf("Exploring %s...\n", name)

	var response pokeapi.PokemonsInLocation
	var err error

	cache := cfg.pokeapiCache
	entry, exist := cache.Get(name)

	if exist {
		err = json.Unmarshal(entry, &response)
	} else {
		response, err = cfg.pokeapiClient.GetPokemonsInLocation(name)
	}

	if err != nil {
		return err
	}

	for _, name := range response {
		fmt.Println(name)
	}

	if !exist {
		b, err := json.Marshal(response)
		if err != nil {
			return err
		}

		cache.Add(name, b)
	}

	return nil
}

func commandCatch(cfg *config, name string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	var pokemonResponse pokeapi.PokemonResponse
	var err error

	cache := cfg.pokeapiCache
	value, exist := cache.Get(name)

	if exist {
		err = json.Unmarshal(value, &pokemonResponse)
	} else {
		pokemonResponse, err = cfg.pokeapiClient.GetPokemonBaseExperience(name)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	dice_roll := r.Intn(100)

	if err != nil {
		return err
	}

	if dice_roll > pokemonResponse.BaseExp {
		fmt.Printf("%s was caught!\n", name)
		cfg.pokedex[name] = pokemonResponse.Pokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	b, err := json.Marshal(pokemonResponse)
	if err != nil {
		return err
	}

	cfg.pokeapiCache.Add(name, b)

	return nil
}

func commandInspect(cfg *config, name string) error {
	pokemon, ok := cfg.pokedex[name]
	if !ok {
		fmt.Printf("You have not caught that pokemon\n")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t-%s: %d\n", stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, typ := range pokemon.Types {
		fmt.Printf("\t-%s\n", typ)
	}
	return nil
}
