package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonBaseExperience(name string) (PokemonResponse, error) {
	url := fmt.Sprintf("%s%s", pokemonURL, name)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return PokemonResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return PokemonResponse{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, res.Body)
	}

	response := pokemonAPIResponse{}
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return PokemonResponse{}, err
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		return PokemonResponse{}, err
	}

	stats := []PokemonStats{}
	types := []string{}

	for _, entry := range response.Stats {
		stats = append(stats, PokemonStats{
			Name:     entry.Stat.Name,
			BaseStat: entry.BaseStat,
		})
	}

	for _, entry := range response.Types {
		types = append(types, entry.Type.Name)
	}

	pokemon := PokemonInfo{
		Name:   response.Name,
		Height: response.Height,
		Weight: response.Weight,
		Stats:  stats,
		Types:  types,
	}

	pokemonResponse := PokemonResponse{
		Pokemon: pokemon,
		BaseExp: response.BaseExperience,
	}

	return pokemonResponse, nil
}
