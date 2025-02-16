package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonsInLocation(locationName string) (PokemonResponse, error) {
	url := fmt.Sprintf("%s%s", exploreLocationURL, locationName)

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

	locations_response := namedLocationAreaAPIResponse{}
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return PokemonResponse{}, err
	}

	err = json.Unmarshal(body, &locations_response)

	if err != nil {
		return PokemonResponse{}, err
	}

	names := []string{}

	for _, entry := range locations_response.PokemonEncounters {
		names = append(names, entry.Pokemon.Name)
	}

	return names, nil
}
