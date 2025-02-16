package main

import (
	"time"

	pokeapi "github.com/EveBisk/pokedex/internal/pokeapi"
	pokecache "github.com/EveBisk/pokedex/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient()
	cache := pokecache.NewCache(5 * time.Second)

	cfg := &config{
		pokeapiClient: pokeClient,
		pokeapiCache:  cache,
		pokedex:       make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
