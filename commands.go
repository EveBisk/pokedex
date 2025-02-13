package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/EveBisk/pokedex/internal/pokeapi"
)

func commandExit(reg *config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)

	return nil
}

func commandHelp(reg *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	registry := getCommands()

	for _, val := range registry {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}

	return nil
}

func commandMap(cfg *config) error {
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

func commandMapBack(cfg *config) error {
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
