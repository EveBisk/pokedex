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

	var response pokeapi.PokemonResponse
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

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	dice_roll := r.Intn(100)

	poke_exp, pokemon, err := cfg.pokeapiClient.GetPokemonBaseExperience(name)

	if err != nil {
		return err
	}

	if dice_roll > poke_exp {
		fmt.Printf("%s was caught!\n", name)
		cfg.pokedex[name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}
