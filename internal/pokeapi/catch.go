package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonBaseExperience(name string) (int, Pokemon, error) {
	url := fmt.Sprintf("%s%s", pokemonURL, name)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return 0, Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return 0, Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return 0, Pokemon{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, res.Body)
	}

	pokemon_response := pokemonAPIResponse{}
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return 0, Pokemon{}, err
	}

	err = json.Unmarshal(body, &pokemon_response)

	if err != nil {
		return 0, Pokemon{}, err
	}

	stats := []PokemonStats{}
	types := []string{}

	for _, entry := range pokemon_response.Stats {
		stats = append(stats, PokemonStats{
			Name:     entry.Stat.Name,
			BaseStat: entry.BaseStat,
		})
	}

	for _, entry := range pokemon_response.Types {
		types = append(types, entry.Type.Name)
	}

	pokemon := Pokemon{
		Name:   pokemon_response.Name,
		Height: pokemon_response.Height,
		Weight: pokemon_response.Weight,
		Stats:  stats,
		Types:  types,
	}

	return pokemon_response.BaseExperience, pokemon, nil
}
