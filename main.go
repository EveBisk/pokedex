package main

import (
	pokeapi "github.com/EveBisk/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient()
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
