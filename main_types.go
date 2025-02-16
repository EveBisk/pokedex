package main

import (
	pokeapi "github.com/EveBisk/pokedex/internal/pokeapi"
	pokecache "github.com/EveBisk/pokedex/internal/pokecache"
)

type config struct {
	pokeapiClient    pokeapi.Client
	pokeapiCache     *pokecache.Cache
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]pokeapi.PokemonInfo
}
