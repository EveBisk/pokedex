package main

import (
	"errors"
	"fmt"
	"os"

	pokeapi "github.com/EveBisk/pokedex/internal/pokeapi"
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

func commandMap(reg *config) error {
	pageUrl := reg.nextLocationsURL

	client := pokeapi.NewClient()
	locations_response, err := client.ListLocations(pageUrl)

	if err != nil {
		return err
	}

	for _, entry := range locations_response.Results {
		fmt.Println(entry.Name)
	}

	reg.prevLocationsURL = locations_response.Previous
	reg.nextLocationsURL = locations_response.Next

	return nil
}

func commandMapBack(reg *config) error {
	if reg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	pageUrl := reg.prevLocationsURL

	client := pokeapi.NewClient()
	locations_response, err := client.ListLocations(pageUrl)

	if err != nil {
		return err
	}

	for _, entry := range locations_response.Results {
		fmt.Println(entry.Name)
	}

	reg.prevLocationsURL = locations_response.Previous
	reg.nextLocationsURL = locations_response.Next

	return nil
}
